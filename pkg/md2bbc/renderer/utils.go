package renderer

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"html"
	"io"
	"sort"
	"strings"
)

var Escaper = [256][]byte{
	'&': []byte("&amp;"),
	'<': []byte("&lt;"),
	'>': []byte("&gt;"),
	'"': []byte("&quot;"),
}

func EscapeHTML(w io.Writer, d []byte) {
	var start, end int
	n := len(d)
	for end < n {
		escSeq := Escaper[d[end]]
		if escSeq != nil {
			w.Write(d[start:end])
			w.Write(escSeq)
			start = end + 1
		}
		end++
	}
	if start < n && end <= n {
		w.Write(d[start:end])
	}
}

func escLink(w io.Writer, text []byte) {
	unesc := html.UnescapeString(string(text))
	EscapeHTML(w, []byte(unesc))
}

func (r *BbcodeRenderer) outOneOfCr(w io.Writer, outFirst bool, first string, second string) {
	if outFirst {
		r.cr(w)
		r.outs(w, first)
	} else {
		r.outs(w, second)
		r.cr(w)
	}
}

func tagWithAttributes(name string, attrs []string) string {
	s := name
	if len(attrs) > 0 {
		s += " " + strings.Join(attrs, " ")
	}
	return s + "]"
}

// BlockAttrs takes a node and checks if it has block level attributes set. If so it
// will return a slice each containing a "key=value(s)" string.
func BlockAttrs(node ast.Node) []string {
	var attr *ast.Attribute
	if c := node.AsContainer(); c != nil && c.Attribute != nil {
		attr = c.Attribute
	}
	if l := node.AsLeaf(); l != nil && l.Attribute != nil {
		attr = l.Attribute
	}
	if attr == nil {
		return nil
	}

	var s []string
	if attr.ID != nil {
		s = append(s, fmt.Sprintf(`%s="%s"`, IDTag, attr.ID))
	}

	classes := ""
	for _, c := range attr.Classes {
		classes += " " + string(c)
	}
	if classes != "" {
		s = append(s, fmt.Sprintf(`class="%s"`, classes[1:])) // skip space we added.
	}

	// sort the attributes so it remain stable between runs
	var keys = []string{}
	for k, _ := range attr.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		s = append(s, fmt.Sprintf(`%s="%s"`, k, attr.Attrs[k]))
	}

	return s
}

func (r *BbcodeRenderer) outTag(w io.Writer, name string, attrs []string) {
	s := name
	if len(attrs) > 0 {
		s += " " + strings.Join(attrs, " ")
	}
	io.WriteString(w, s+">")
	r.lastOutputLen = 1
}

func isListItem(node ast.Node) bool {
	_, ok := node.(*ast.ListItem)
	return ok
}

func (r *BbcodeRenderer) outs(w io.Writer, s string) {
	r.lastOutputLen = len(s)
	io.WriteString(w, s)
}

func (r *BbcodeRenderer) out(w io.Writer, d []byte) {
	r.lastOutputLen = len(d)
	w.Write(d)
}

func (r *BbcodeRenderer) outOneOf(w io.Writer, outFirst bool, first string, second string) {
	if outFirst {
		r.outs(w, first)
	} else {
		r.outs(w, second)
	}
}

func (r *BbcodeRenderer) cr(w io.Writer) {
	if r.lastOutputLen > 0 {
		r.outs(w, "\n")
	}
}

func (r *BbcodeRenderer) hardBreak(w io.Writer, node *ast.Hardbreak) {
	r.outs(w, "\n")
	r.cr(w)
}
