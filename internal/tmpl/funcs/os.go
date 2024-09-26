package funcs

import (
	"bytes"
	"io"
	"net/http"
	"text/template"
)

func TxtFrom(srcs ...string) (string, error) {
	buf := bytes.NewBuffer(nil)
	for _, src := range srcs {
		resp, err := http.Get(src)
		if err != nil {
			return "", err
		}
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		buf.Write(body)
		_ = resp.Body.Close()
	}
	return buf.String(), nil
}

// Tmpl 注意，此template不提供額外的FuncMap
func Tmpl(src string, ctx map[string]any) (string, error) {
	resp, err := http.Get(src)
	if err != nil {
		return "", err
	}
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_ = resp.Body.Close()

	var t *template.Template
	t, err = template.New("").Parse(string(body))
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err = t.Execute(buf, ctx); err != nil {
		return "", err
	}
	return buf.String(), nil
}
