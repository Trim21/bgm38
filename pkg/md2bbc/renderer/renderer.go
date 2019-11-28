package renderer

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"io"
)

var IDTag = "id"

type BbcodeRenderer struct {
	markdown.Renderer
	lastOutputLen int
}

func (BbcodeRenderer) RenderHeader(w io.Writer, ast ast.Node) {
}
func (BbcodeRenderer) RenderFooter(w io.Writer, ast ast.Node) {
}

func (r BbcodeRenderer) text(w io.Writer, text *ast.Text) {
	_, parentIsLink := text.Parent.(*ast.Link)
	if parentIsLink {
		return
		//escLink(w, text.Literal)
	}

	_, parentIsImg := text.Parent.(*ast.Image)
	if parentIsImg {
		return
	}
	r.out(w, text.Literal)
}

func (r *BbcodeRenderer) paragraphEnter(w io.Writer, para *ast.Paragraph) {
	// TODO: untangle this clusterfuck about when the newlines need
	// to be added and when not.
	prev := ast.GetPrevNode(para)
	if prev != nil {
		switch prev.(type) {
		case *ast.HTMLBlock, *ast.List, *ast.Paragraph, *ast.Heading, *ast.CaptionFigure, *ast.CodeBlock, *ast.BlockQuote, *ast.Aside, *ast.HorizontalRule:
			r.cr(w)
		}
	}

	if prev == nil {
		_, isParentBlockQuote := para.Parent.(*ast.BlockQuote)
		if isParentBlockQuote {
			r.cr(w)
		}
		_, isParentAside := para.Parent.(*ast.Aside)
		if isParentAside {
			r.cr(w)
		}
	}

	tag := tagWithAttributes("<p", BlockAttrs(para))
	r.outs(w, tag)
}

func (r *BbcodeRenderer) paragraphExit(w io.Writer, para *ast.Paragraph) {
	r.outs(w, "</p>")
	if !(isListItem(para.Parent) && ast.GetNextNode(para) == nil) {
		r.cr(w)
	}
}

func (r *BbcodeRenderer) paragraph(w io.Writer, para *ast.Paragraph, entering bool) {
	if entering {
		r.paragraphEnter(w, para)
	} else {
		r.paragraphExit(w, para)
	}
}

func (r BbcodeRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Text:
		r.text(w, node)
	case *ast.Softbreak:
		r.cr(w)
	case *ast.Hardbreak:
		r.hardBreak(w, node)
	case *ast.Emph:
		r.outOneOf(w, entering, "[i]", "[/i]")
	case *ast.Strong:
		r.outOneOf(w, entering, "[b]", "[/b]")
	case *ast.Del:
		r.outOneOf(w, entering, "[s]", "[/s]")
	case *ast.BlockQuote:
		//tag := tagWithAttributes("[quote", BlockAttrs(node))
		//r.outs(w, "[quote]")
		r.outOneOfCr(w, entering, "[quote]\n", "\n[/quote]")
	case *ast.Link:
		r.link(w, node, entering)
	case *ast.Image:
		r.image(w, node, entering)
	case *ast.Code:
		r.code(w, node)
	case *ast.CodeBlock:
		r.codeBlock(w, node)
	//case *ast.Aside:
	//	tag := tagWithAttributes("<aside", BlockAttrs(node))
	//	r.outOneOfCr(w, entering, tag, "</aside>")
	//case *ast.CrossReference:
	//	link := &ast.Link{Destination: append([]byte("#"), node.Destination...)}
	//	r.link(w, link, entering)
	//case *ast.Citation:
	//	r.citation(w, node)

	//case *ast.Caption:
	//	r.caption(w, node, entering)
	//case *ast.CaptionFigure:
	//	r.captionFigure(w, node, entering)
	//case *ast.Document:
	//do
	//nothing
	case *ast.Paragraph:
		//r.paragraph(w, node, entering)
		//break
		//r.cr(w)
	//case *ast.HTMLSpan:
	//	r.htmlSpan(w, node)
	//case *ast.HTMLBlock:
	//	r.htmlBlock(w, node)
	case *ast.Heading:
		r.heading(w, node, entering)
		//case *ast.HorizontalRule:
		//	r.horizontalRule(w, node)
		//case *ast.List:
		//	r.list(w, node, entering)
		//case *ast.ListItem:
		//	r.listItem(w, node, entering)
		//case *ast.Table:
		//	tag := tagWithAttributes("<table", BlockAttrs(node))
		//	r.outOneOfCr(w, entering, tag, "</table>")
		//case *ast.TableCell:
		//	r.tableCell(w, node, entering)
		//case *ast.TableHeader:
		//	r.outOneOfCr(w, entering, "<thead>", "</thead>")
		//case *ast.TableBody:
		//	r.tableBody(w, node, entering)
		//case *ast.TableRow:
		//	r.outOneOfCr(w, entering, "<tr>", "</tr>")
		//case *ast.TableFooter:
		//	r.outOneOfCr(w, entering, "<tfoot>", "</tfoot>")
		//case *ast.Math:
		//	r.outOneOf(w, true, `<span class="math inline">\(`, `\)</span>`)
		//	EscapeHTML(w, node.Literal)
		//	r.outOneOf(w, false, `<span class="math inline">\(`, `\)</span>`)
		//case *ast.MathBlock:
		//	r.outOneOf(w, entering, `<p><span class="math display">\[`, `\]</span></p>`)
		//	if entering {
		//		EscapeHTML(w, node.Literal)
		//	}
		//case *ast.DocumentMatter:
		//	r.matter(w, node, entering)
		//case *ast.Callout:
		//	r.callout(w, node)
		//case *ast.Index:
		//	r.index(w, node)
		//case *ast.Subscript:
		//	r.outOneOf(w, true, "<sub>", "</sub>")
		//	if entering {
		//		Escape(w, node.Literal)
		//	}
		//	r.outOneOf(w, false, "<sub>", "</sub>")
		//case *ast.Superscript:
		//	r.outOneOf(w, true, "<sup>", "</sup>")
		//	if entering {
		//		Escape(w, node.Literal)
		//	}
		//	r.outOneOf(w, false, "<sup>", "</sup>")
		//case *ast.Footnotes:
		//	nothing by default; just output the list.
		//default:
		//	panic(fmt.Sprintf("Unknown node %T", node))
	}
	return ast.GoToNext
}
