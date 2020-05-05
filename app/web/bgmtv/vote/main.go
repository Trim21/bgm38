package vote

import (
	"io"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gofiber/fiber"
	"github.com/wcharczuk/go-chart"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/app/web/res"
	"bgm38/app/web/utils/handler"
	"bgm38/pkg/fetch"
	"bgm38/pkg/parser"
)

func Group(app *fiber.Group) {
	app.Get("/vote/:topic_id", handler.LogError(vote))
}

// @ID voteResult
// @Summary 解析帖子，生成投票结果
// @Description 解析帖子，生成投票结果
// @Produce  image/svg+xml
// @Param topic_id path int true "topic_id, 小组讨论贴的主题"
// @Success 200 {string} string "image/xvg+xml"
// @Failure 401 {object} res.Error
// @Failure 502 {object} res.Error
// @Router /bgm.tv/v1/vote/{topic_id} [get]
func vote(c *fiber.Ctx, logger *zap.Logger) error {
	var topicInput = c.Params("topic_id")
	topicID, err := strconv.Atoi(topicInput)
	if err != nil {
		return c.Status(401).JSON(res.Error{
			Status:  "error",
			Message: "`topic_id` should be int",
		})
	}
	t, err := load(topicID)
	if err != nil {
		return c.Status(502).JSON(res.Error{
			Status:  "error",
			Message: "can't fetch topic html or can't parse it",
		})
	}
	voteOptions, err := extraRawOption(t)
	if err != nil || voteOptions == "" {
		return c.Status(401).JSON(res.Error{
			Status:  "error",
			Message: "can't find valid options, should start with `[code]vote: true`",
		})
	}
	o, err := parseOption(voteOptions)
	if err != nil {
		return c.Status(401).JSON(res.Error{
			Status:  "error",
			Message: "options is not correct",
		})
	}
	c.Fasthttp.Response.Header.SetContentType("image/svg+xml")
	return render(c.Fasthttp.Response.BodyWriter(), t, o)
}

func getVote(doc *html.Node, voteOptionsLen int) (s []int) {
	el := htmlquery.FindOne(doc, ".//div[@class='codeHighlight']/pre")
	if el == nil {
		return
	}
	text := strings.Trim(htmlquery.InnerText(el), "\u00A0 \n")
	if strings.HasPrefix(text, "$") && strings.HasSuffix(text, "$") {
		o := text[1 : len(text)-1]
		v, err := strconv.Atoi(o)
		if err != nil {
			return
		}
		if 0 < v && v <= voteOptionsLen {
			s = append(s, v)
		}
	}
	return
}
func countVotes(userVotes map[string][]int) map[int]int {
	counter := make(map[int]int)
	for _, ints := range userVotes {
		for _, i := range ints {
			if _, ok := counter[i]; !ok {
				counter[i] = 0
			}
			counter[i]++
		}
	}
	return counter
}

func render(w io.Writer, t parser.T, o Options) error {
	var l = o.Len()
	var userVotes = make(map[string][]int)
	for _, reply := range t.Replies {
		v := getVote(reply.RawContent, l)
		if len(v) > 0 {
			userVotes[reply.Author] = v
		}
		for _, subReply := range reply.Replies {
			v := getVote(subReply.RawContent, l)
			if len(v) > 0 {
				userVotes[subReply.Author] = v
			}
		}
	}
	if len(userVotes) == 0 {
		_, err := w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="1" height="1"/>`))
		return err
	}
	var counter = countVotes(userVotes)
	var v chart.Values
	for key, value := range counter {
		v = append(v,
			chart.Value{
				Label: o.Options[key-1],
				Value: float64(value),
			})

	}
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: v,
	}
	return pie.Render(chart.SVG, w)
}

func load(topicID int) (parser.T, error) {
	var t parser.T
	doc, err := fetch.Topic(topicID)
	if err != nil {
		return t, err
	}

	t, err = parser.Topic(doc)
	if err != nil {
		return t, err
	}
	return t, err
}
