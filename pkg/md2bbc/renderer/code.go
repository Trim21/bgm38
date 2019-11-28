package renderer

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"io"
)

func (r *BbcodeRenderer) code(w io.Writer, node *ast.Code) {
	r.outs(w, fmt.Sprintf("[code]%s[/code]", node.Literal))
}

func (r *BbcodeRenderer) codeBlock(w io.Writer, codeBlock *ast.CodeBlock) {
	r.outs(w, fmt.Sprintf("\n[code]\n%s\n[/code]\n", codeBlock.Literal))
}
