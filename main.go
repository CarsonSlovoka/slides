package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	http2 "github.com/CarsonSlovoka/slides/internal/http"
	"github.com/CarsonSlovoka/slides/internal/markdown"
	htmlTemplate "html/template"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
)

//go:embed reveal.js/dist/reveal.js
//go:embed reveal.js/dist/reveal.css
//go:embed reveal.js/dist/reset.css
//go:embed reveal.js/dist/theme/*css
//go:embed reveal.js/plugin/**/*.js
var revealFS embed.FS

//go:embed help.md
var helpFS []byte

func HandleHelp(w http.ResponseWriter, r *http.Request) {
	md := markdown.New()
	tmpl, err := htmlTemplate.New("").Funcs(map[string]any{
		"unsafeHTML": func(s string) htmlTemplate.HTML { return htmlTemplate.HTML(s) },
	}).Parse("{{unsafeHTML .MD}}")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mdBuf := bytes.NewBuffer(nil)
	if err = md.Convert(helpFS, mdBuf); err != nil {
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

func HandleListMD(w http.ResponseWriter, r *http.Request) {
	// todo 顯示 md目錄之中的所有md檔案
}

// HandleTxt 回傳md目錄下的檔案內容，視為純文本
func HandleTxt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.StripPrefix("/txt/", http.FileServer(http.Dir("md"))).ServeHTTP(w, r)
}

func HandleMD(w http.ResponseWriter, r *http.Request) {
	if filepath.Ext(r.URL.Path) != ".md" {
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

	tmpl, err := htmlTemplate.New("slides.gohtml").Funcs(map[string]any{
		"unsafeHTML": func(s string) htmlTemplate.HTML { return htmlTemplate.HTML(s) },
	}).ParseFiles("slides.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, map[string]any{
		"Title":     query.Get("title"),
		"MDPath":    "/txt/" + r.PathValue("mdPath"), // 我們讓其導向 HandleTxt
		"Theme":     theme,
		"View":      view,
		"AutoSlide": autoSlide,
	}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	var enableTls bool
	var port int
	flag.BoolVar(&enableTls, "tls", false, "Enable TLS")
	flag.IntVar(&port, "port", 8080, "port number")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("GET /reveal.js/", http.FileServerFS(revealFS))
	mux.HandleFunc("GET /help", HandleHelp)
	mux.HandleFunc("GET /md", HandleListMD)
	mux.HandleFunc("GET /txt/", HandleTxt)
	mux.HandleFunc("GET /md/{mdPath...}", HandleMD)
	mux.Handle("GET /assets/", http.FileServer(http.Dir(".")))
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
