package fetch

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
)

var client = resty.New()

func Topic(id int) (*html.Node, error) {
	res, err := client.R().SetDoNotParseResponse(true).Get(fmt.Sprintf("https://mirror.bgm.rin.cat/group/topic/%d", id))
	if err != nil {
		return nil, err
	}
	return html.Parse(res.RawBody())
}
