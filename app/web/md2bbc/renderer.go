// Package markdown provides a Markdown renderer.
package md2bbc

import (
	"bytes"
	"fmt"

	"github.com/russross/blackfriday"
)

type markdownRenderer struct {
	blackfriday.Html
	normalTextMarker   map[*bytes.Buffer]int
	orderedListCounter map[int]int
	paragraph          map[int]bool // Used to keep track of whether a given list item uses a paragraph for large spacing.
	listDepth          int
	lastNormalText     string
}

// Block-level callbacks.
func (*markdownRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	doubleSpace(out)
	out.WriteString("[code]")
	out.WriteString("\n")
	out.Write(text)
	out.WriteString("[/code]\n")
}

func (*markdownRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	doubleSpace(out)
	out.WriteString("[quote]\n")
	lines := bytes.Split(text, []byte("\n"))
	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}
		if len(line) != 0 {
			out.Write(line)
		}
		out.WriteString("\n")
	}
	out.WriteString("[/quote]\n")
}

func (mr *markdownRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()
	doubleSpace(out)

	out.WriteString(fmt.Sprintf("[size=%d]", 26-level*2))

	if !text() {
		out.Truncate(marker)
		fmt.Println("Truncate")
		return
	}

	// if !text() {
	// 	out.Truncate(textMarker)
	// 	return
	// }
	// out.Truncate(out.Len() - len([]byte(out.String()[textMarker:])))
	// out.WriteString("\n")
	out.WriteString("[/size]\n")
	// out.WriteString("\n")

}

func (*markdownRenderer) HRule(out *bytes.Buffer) {
	doubleSpace(out)
	out.WriteString("---\n")
}

func (mr *markdownRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	marker := out.Len()
	doubleSpace(out)

	mr.listDepth++
	defer func() { mr.listDepth-- }()
	if flags&blackfriday.LIST_TYPE_ORDERED != 0 {
		mr.orderedListCounter[mr.listDepth] = 1
	}
	if !text() {
		out.Truncate(marker)
		return
	}
}

func (mr *markdownRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	if flags&blackfriday.LIST_TYPE_ORDERED != 0 {
		fmt.Fprintf(out, "%d.", mr.orderedListCounter[mr.listDepth])
		out.WriteByte(' ')
		out.Write(text)
		mr.orderedListCounter[mr.listDepth]++
	} else {
		out.WriteString("-")
		out.WriteByte(' ')
		out.Write(text)
	}
	out.WriteString("\n")
	if mr.paragraph[mr.listDepth] {
		if flags&blackfriday.LIST_ITEM_END_OF_LIST == 0 {
			out.WriteString("\n")
		}
		mr.paragraph[mr.listDepth] = false
	}
}

func (mr *markdownRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()
	doubleSpace(out)

	mr.paragraph[mr.listDepth] = true

	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("\n")
}

func (mr *markdownRenderer) Table(out *bytes.Buffer, header, body []byte, columnData []int) {
	out.WriteString("\n<Not Supported In BBCode.>\n")
}

func (*markdownRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	out.WriteString("<Footnotes: Not implemented.>")
}
func (*markdownRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	out.WriteString("<FootnoteItem: Not implemented.>")
}

// Span-level callbacks.
func (mr *markdownRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	mr.Link(out, link, nil, nil)
}
func (*markdownRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	out.WriteString("[code]")
	out.Write(text)
	out.WriteString("[/code]")
}
func (mr *markdownRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("[b]")
	out.Write(text)
	out.WriteString("[/b]")
}
func (*markdownRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	if len(text) == 0 {
		return
	}
	out.WriteString("[i]")
	out.Write(text)
	out.WriteString("[/i]")
}
func (*markdownRenderer) Image(out *bytes.Buffer, link, title, alt []byte) {
	out.WriteString("[img]")
	out.Write(escape(link))
	out.WriteString("[/img]")
}

func (*markdownRenderer) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (*markdownRenderer) Link(out *bytes.Buffer, link, title, content []byte) {
	if len(content) > 0 {
		out.WriteString("[url=")
		out.Write(escape(link))
		out.WriteString("]")
		out.Write(content)
	} else {
		out.WriteString("[url]")
		out.Write(escape(link))
	}
	out.WriteString("[/url]")
}
func (*markdownRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("[b][i]")
	out.Write(text)
	out.WriteString("[/i][/b]")
}
func (*markdownRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.WriteString("~~")
	out.Write(text)
	out.WriteString("~~")
}
func (*markdownRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	out.WriteString("<FootnoteRef: Not implemented.>")
}

// escape replaces instances of backslash with escaped backslash in text.
func escape(text []byte) []byte {
	return bytes.Replace(text, []byte(`\`), []byte(`\\`), -1)
}

func isNumber(data []byte) bool {
	for _, b := range data {
		if b < '0' || b > '9' {
			return false
		}
	}
	return len(data) > 0
}

func needsEscaping(text []byte, lastNormalText string) bool {
	switch string(text) {
	case `\`,
		"`",
		"*",
		"_",
		"{", "}",
		"[", "]",
		"(", ")",
		"#",
		"+",
		"-":
		return true
	case "!":
		return false
	case ".":
		// Return true if number, because a period after a number must be escaped to not get parsed as an ordered list.
		return isNumber([]byte(lastNormalText))
	case "<", ">":
		return true
	default:
		return false
	}
}

// Low-level callbacks.
func (*markdownRenderer) Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}
func (mr *markdownRenderer) NormalText(out *bytes.Buffer, text []byte) {
	normalText := string(text)
	if needsEscaping(text, mr.lastNormalText) {
		text = append([]byte("\\"), text...)
	}
	mr.lastNormalText = normalText
	if mr.listDepth > 0 && string(text) == "\n" {
		return
	}
	cleanString := cleanWithoutTrim(string(text))
	if cleanString == "" {
		return
	}
	if mr.skipSpaceIfNeededNormalText(out, cleanString) {
		// Skip first space if last character is already a space (i.e., no need for a 2nd space in a row).
		cleanString = cleanString[1:]
	}
	out.WriteString(cleanString)
	if len(cleanString) >= 1 && cleanString[len(cleanString)-1] == ' ' {
		// If it ends with a space, make note of that.
		mr.normalTextMarker[out] = out.Len()
	}
}

func (mr *markdownRenderer) skipSpaceIfNeededNormalText(out *bytes.Buffer, cleanString string) bool {
	if cleanString[0] != ' ' {
		return false
	}
	if _, ok := mr.normalTextMarker[out]; !ok {
		mr.normalTextMarker[out] = -1
	}
	return mr.normalTextMarker[out] == out.Len()
}

// cleanWithoutTrim is like clean, but doesn't trim blanks.
func cleanWithoutTrim(s string) string {
	var b []byte
	var p byte
	for i := 0; i < len(s); i++ {
		q := s[i]
		if q == '\n' || q == '\r' || q == '\t' {
			q = ' '
		}
		if q != ' ' || p != ' ' {
			b = append(b, q)
			p = q
		}
	}
	return string(b)
}

func doubleSpace(out *bytes.Buffer) {
	if out.Len() > 0 {
		out.WriteByte('\n')
	}
}

// NewRenderer returns a Markdown renderer.
// If opt is nil the defaults are used.
func NewRenderer() blackfriday.Renderer {
	mr := &markdownRenderer{
		normalTextMarker:   make(map[*bytes.Buffer]int),
		orderedListCounter: make(map[int]int),
		paragraph:          make(map[int]bool),
	}
	return mr
}
