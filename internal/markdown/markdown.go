package markdown

import (
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func New() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // 包含 Linkify, Table, Strikethrough, TaskList
			extension.Footnote,

			meta.Meta, // front-matter toml格式
			highlighting.Highlighting,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // 會自動加上id，但如果有中文heading會不支持

			// https://github.com/yuin/goldmark#attributes
			parser.WithAttribute(), // 推薦補上這個，可以在heading旁邊利用## MyH1{class="cls1 cls2"}來補上一些屬性 // https://www.markdownguide.org/extended-syntax/#heading-ids // https://github.com/gohugoio/hugo/issues/7548
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
}
