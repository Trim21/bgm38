package renderer

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown/ast"
)

func (r *BbcodeRenderer) image(w io.Writer, node *ast.Image, entering bool) {
	if entering {
		r.imageEnter(w, node)
	} else {
		r.imageExit(w, node)
	}
}

func (r *BbcodeRenderer) imageEnter(w io.Writer, image *ast.Image) {
	dest := image.Destination
	fmt.Fprintf(w, "[img]%s", dest)
	//r.outs(w, "[img]"+dest)
	//escLink(w, dest)
}

func (r *BbcodeRenderer) imageExit(w io.Writer, image *ast.Image) {
	r.outs(w, `[/img]`)
}
