package vote

import (
	"strings"

	"github.com/antchfx/htmlquery"

	"bgm38/pkg/parser"
)

func extraRawOption(t parser.T) (voteOptions string, err error) {
	codeBlocks, err := htmlquery.QueryAll(t.RawContent, ".//pre")

	if err != nil {
		return
	}
	for _, el := range codeBlocks {
		content := htmlquery.OutputHTML(el, false)
		if strings.HasPrefix(content, "vote: true") || strings.HasPrefix(content, "vote:true") {
			voteOptions = strings.ReplaceAll(content, "<br/>", "")
		}
	}
	return
}
