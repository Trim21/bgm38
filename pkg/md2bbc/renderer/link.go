package renderer

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"io"
)

func (r *BbcodeRenderer) linkEnter(w io.Writer, link *ast.Link) {
	dest := link.Destination

	if len(link.Title) > 0 {
		fmt.Fprintf(w, "[url=%s]%s", dest, link.Title)
	} else {
		fmt.Fprintf(w, "[url]%s", dest)
	}

}

func (r *BbcodeRenderer) linkExit(w io.Writer, link *ast.Link) {
	if link.NoteID == 0 {
		r.outs(w, "[/url]")
	}
}

func (r *BbcodeRenderer) link(w io.Writer, link *ast.Link, entering bool) {
	if entering {
		r.linkEnter(w, link)
	} else {
		r.linkExit(w, link)
	}
}
