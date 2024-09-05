package markdown_test

import (
	"bytes"
	"github.com/CarsonSlovoka/slides/internal/markdown"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"strings"
	"testing"
)

type emptyWriter struct{}

func (w *emptyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func TestNew(t *testing.T) {
	m := markdown.New()
	out := bytes.NewBuffer(nil)
	content := []byte(`---
Title: goldmark-meta
Summary: Add YAML metadata to the document
Tags:
    - markdown
    - goldmark
Langs: [py, javascript]
---

# Hello goldmark-meta
`)
	ctx := parser.NewContext() // 你可以得到frontMatter的內容
	// m.Convert(content, nil, parser.WithContext(ctx)) // panic, 一定要給對象
	// m.Convert(content, &emptyWriter{}, parser.WithContext(ctx)) // 錯誤，會遇到short write的錯誤
	if err := m.Convert(content, out, parser.WithContext(ctx)); err != nil {
		t.Fatal(err)
	}

	var metaData map[string]any
	metaData = meta.Get(ctx)
	if metaData["notExists"] != nil { // map取不到(即便map還是nil)的時候返回元素的初始值 https://go.dev/play/p/g3hUOvOv9iO
		t.Fatal()
	}
	if l := metaData["Langs"].([]any); len(l) != 2 || l[0] != "py" {
		t.Fatal()
	}

	for _, test := range []struct {
		actual   any
		expected any
	}{
		{metaData["Title"].(string), "goldmark-meta"},
		{len(metaData["Tags"].([]any)), 2},
		{metaData["Tags"].([]any)[0], "markdown"},
		{strings.Contains(string(out.Bytes()), `<h1 id="hello-goldmark-meta">Hello goldmark-meta</h1>`), true},
	} {
		if test.actual != test.expected {
			t.Fatalf("acutal: %+v\nexpected: %+v", test.actual, test.expected)
		}
	}
}

func TestNew2(t *testing.T) {
	m := markdown.New()
	out := bytes.NewBuffer(nil)
	content := []byte(`# Hello`)
	ctx := parser.NewContext()
	if err := m.Convert(content, out, parser.WithContext(ctx)); err != nil {
		t.Fatal(err)
	}
	if meta.Get(ctx) != nil {
		// 沒有任何的frontMatter，得到的就是nil
		t.Fatal()
	}
	fm := meta.Get(ctx)    // 這裡的nil指的是map[string]interface{}的對象，而非any
	if fm["Tags"] != nil { // 因此以下的判斷不會觸發panic
		t.Fatal()
	}
}
