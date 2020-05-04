package parser

import (
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type SubReply struct {
	Time   time.Time
	Author string
	Floor  string
}

type Reply struct {
	Time   time.Time
	Author string
	Floor  string
}

type T struct {
	Title  string
	Author string
	Time   time.Time
}

func Topic(doc *html.Node) (T, error) {
	var t = T{}
	t.Title = htmlquery.InnerText(htmlquery.FindOne(doc, "/html/head/title"))
	topicMain := htmlquery.FindOne(doc, "//*[contains(@class, 'postTopic')]")
	t.Author = htmlquery.SelectAttr(htmlquery.FindOne(topicMain, "a[@class='avatar']"), "href")
	return t, nil
}

func reInfo(s string) (floor string, t time.Time, err error) {
	// #1 - 2020-5-4 04:27
	l := strings.SplitN(s, " - ", 2)
	floor = l[0]
	t, err = time.Parse("2006-1-2 15:04", l[1])
	return
}
