package renderer

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"io"
)

func headingOpenTagFromLevel(level int) string {
	return fmt.Sprintf("[size=%d]", 26-level*2)
}

func headingCloseTagFromLevel(level int) string {
	return "[/size]"
}

func (r *BbcodeRenderer) headingEnter(w io.Writer, nodeData *ast.Heading) {
	//r.cr(w)
	r.outs(w, headingOpenTagFromLevel(nodeData.Level))
	//r.outs(w, string(nodeData.Literal))
	//r.outs(w, "[hE]")
	//r.cr(w)
}

func (r *BbcodeRenderer) headingExit(w io.Writer, heading *ast.Heading) {
	r.outs(w, headingCloseTagFromLevel(heading.Level))
	//if !(isListItem(heading.Parent) && ast.GetNextNode(heading) == nil) {
	//r.outs(w, "[hX]")
	r.cr(w)
	//}
}

func (r *BbcodeRenderer) heading(w io.Writer, node *ast.Heading, entering bool) {
	if entering {
		r.headingEnter(w, node)
	} else {
		r.headingExit(w, node)
	}
}
