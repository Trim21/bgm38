package md2bbc

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

var renderer *html.Renderer

func init() {
}

func markdown_to_bbcode(rawMarkdown []byte) string {
	md := []byte("markdown text")
	bbcode := markdown.ToHTML(md, nil, renderer)
	return string(bbcode)
}
