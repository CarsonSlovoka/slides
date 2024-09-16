package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	http2 "github.com/CarsonSlovoka/slides/internal/http"
	"github.com/CarsonSlovoka/slides/internal/markdown"
	"github.com/CarsonSlovoka/slides/internal/tmpl/funcs"
	htmlTemplate "html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	textTemplate "text/template"
)

//go:embed reveal.js/dist/reveal.js
//go:embed reveal.js/dist/reveal.css
//go:embed reveal.js/dist/reset.css
//go:embed reveal.js/dist/theme/*css
//go:embed reveal.js/dist/theme/fonts/*
//go:embed reveal.js/plugin/**/*.js
//go:embed reveal.js/plugin/**/*.css
var revealFS embed.FS

//go:embed plugin/**/*.css
//go:embed plugin/**/*.mjs
//go:embed plugin/**/*.js
var pluginFS embed.FS

//go:embed help.md
var helpBytes []byte

func HandleHelp(w http.ResponseWriter, r *http.Request) {
	md := markdown.New()
	tmpl, err := htmlTemplate.New("").Funcs(map[string]any{
		"unsafeHTML": func(s string) htmlTemplate.HTML { return htmlTemplate.HTML(s) },
	}).Parse(
		`
<head><link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"></head>
<main class="container">{{unsafeHTML .MD}}</main>
`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 將 helpBytes 中的 context 做置換
	helpGoHtml, err := textTemplate.New("").Funcs(nil).Parse(string(helpBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resolveContent := bytes.NewBuffer(nil)
	if err = helpGoHtml.Execute(resolveContent, map[string]string{
		"MDName": mdFolderName,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mdBuf := bytes.NewBuffer(nil)
	if err = md.Convert(resolveContent.Bytes(), mdBuf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, map[string]any{
		"MD": mdBuf.String(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandleListMD 查看所有能投影的清單 (也就是位於./md目錄之中所有的md檔案)
func HandleListMD(w http.ResponseWriter, r *http.Request) {
	html := `<head><link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"></head>
<main class=container><ul>
`
	if err := filepath.Walk(fmt.Sprintf("./%s", mdFolderName), func(p string, info fs.FileInfo, err error) error {
		if info == nil {
			return fmt.Errorf("file %s does not exist, 請確定你的工作路徑是否正確", p)
		}

		if info.IsDir() || filepath.Ext(strings.ToLower(p)) != ".md" {
			return nil
		}

		html += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n",
			filepath.ToSlash(p), // convert to forward slash
			p,
		)
		return nil
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html += "</ul></main>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprintf(w, html)
}

// HandleTxt 回傳md目錄下的檔案內容，視為純文本
func HandleTxt(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query["raw"] != nil {
		// 直接返回文本內容
		w.Header().Set("Content-Type", "text/plain")
		http.StripPrefix("/txt/", http.FileServer(http.Dir(mdFolderName))).ServeHTTP(w, r)
		return
	}

	mdPath := filepath.Join(mdFolderName, r.URL.Path[5:])
	bs, err := os.ReadFile(mdPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 用go template處理之後再返回
	maps := funcs.Maps{
		"dict": funcs.Dict,
		"list": funcs.List, // slice 已經是保留字了，所以用list
	}.AddMaps(funcs.MathMaps()).Add(
		"unsafeHTML", func(s string) htmlTemplate.HTML { return htmlTemplate.HTML(s) },
	)
	tmpl := htmlTemplate.New("").Funcs(maps)
	tmpl, err = tmpl.Parse(string(bs))
	if err != nil {
		log.Printf("%s: %s\n", mdPath, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleMD(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(filepath.Ext(r.URL.Path)) != ".md" {
		// 使如果你將圖片放在md的目錄中，也能夠顯示
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
		return
	}

	query := r.URL.Query()
	var theme string
	if query["theme"] != nil {
		// https://revealjs.com/themes/
		theme = query.Get("theme")
	}
	if theme == "" {
		theme = "black"
	}

	view := query.Get("view")

	var autoSlide uint64
	if millSec := query.Get("autoSlide"); millSec != "" {
		if view == "scroll" {
			log.Println("autoSlide not working on view=scroll")
		}
		var err error
		autoSlide, err = strconv.ParseUint(millSec, 10, 64)
		if err != nil {
			log.Println(err)
			autoSlide = 0
		}
	}
	tmpl := htmlTemplate.New("slides.gohtml").Funcs(map[string]any{
		"unsafeHTML": func(s string) htmlTemplate.HTML { return htmlTemplate.HTML(s) },
	})
	var err error
	if _, err = os.Stat("slides.gohtml"); err == nil {
		// 我們選擇直接讀檔的方式，這樣可以提供即時修改也不需要再重新啟動程式
		tmpl, err = tmpl.ParseFiles("slides.gohtml")
	} else {
		tmpl, err = tmpl.Parse(string(slidesTmplBytes))
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mdPath := "/txt/" + r.PathValue("mdPath") // 我們讓其導向 HandleTxt
	if query["raw"] != nil {
		mdPath += "?raw"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err = tmpl.Execute(w, map[string]any{
		"Title":     query.Get("title"),
		"MDPath":    mdPath,
		"Theme":     theme,
		"View":      view,
		"AutoSlide": autoSlide,
	}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func init() {
	if _, err := os.Stat("slides.gohtml"); !os.IsNotExist(err) {
		// 如果找尋的到使用者自己提供的slides.gohtml模板內容就已使用者提供的為準
		if slidesTmplBytes != nil {
			log.Println("overwrite slidesTmplBytes")
		}
		slidesTmplBytes, err = os.ReadFile("slides.gohtml")
		if err != nil {
			log.Fatal(err)
		}
	}

	// 如果你的打包預設有將模板匯入go build -tags tmpl
	// 那麼 slidesTmplBytes 可以有預設的嵌入內容
	if slidesTmplBytes == nil {
		log.Fatal("you must provide `slides.gohtml`")
	}
}

var mdFolderName string

func main() {
	var enableTls bool
	var port int

	flag.BoolVar(&enableTls, "tls", false, "Enable TLS")
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&mdFolderName, "md", "md", "md folder")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("GET /reveal.js/", http.FileServerFS(revealFS))
	mux.HandleFunc("GET /help", HandleHelp)
	mux.HandleFunc("GET /txt/", HandleTxt)
	mux.HandleFunc(fmt.Sprintf("GET /%s", mdFolderName), HandleListMD)
	mux.HandleFunc(fmt.Sprintf("GET /%s/{mdPath...}", mdFolderName), HandleMD)
	mux.Handle("GET /assets/", http.FileServer(http.Dir(".")))

	// 我們內嵌的plugin
	mux.Handle("GET /slides/plugin/", http.StripPrefix("/slides/", http.FileServerFS(pluginFS)))
	// 讓使用者有辦法載入自己的plugin
	mux.Handle("GET /plugin/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/", HandleHelp)

	fmt.Println("site routes:")
	patterns := reflect.ValueOf(mux).Elem().FieldByName("patterns")
	for i := 0; i < patterns.Len(); i++ {
		pattern := patterns.Index(i).Elem()
		strField := pattern.FieldByName("str").String()
		fmt.Println(strField)
	}

	go func() {
		log.Printf("Server is running on http://127.0.0.1:%d\n", port)
		if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux); err != nil {
			log.Fatal(err)
		}
	}()

	if enableTls {
		go func() {
			log.Printf("Server is running on https://localhost:%d\n", port)
			if err := http2.ListenAndServeTLS(mux, fmt.Sprintf(":%d", port)); err != nil {
				log.Println(err)
			}
		}()
	}

	select {}
}
