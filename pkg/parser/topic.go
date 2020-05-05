package parser

import (
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Floor [2]int

type SubReply struct {
	Time   time.Time
	Author string
	Floor  Floor
}

type Reply struct {
	Time    time.Time
	Author  string
	Floor   Floor
	Replies []SubReply
}

type T struct {
	Title      string
	Author     string
	Time       time.Time
	RawContent *html.Node
	Replies    []Reply
}

func Topic(doc *html.Node) (t T, err error) {
	t.Title = htmlquery.InnerText(htmlquery.FindOne(doc, "/html/head/title/text()"))
	topicMain := htmlquery.FindOne(doc, "//*[contains(@class, 'postTopic')]")
	t.Author = htmlquery.InnerText(htmlquery.FindOne(topicMain, "a[@class='avatar']/@href"))
	t.RawContent = htmlquery.FindOne(topicMain, ".//div[@class='topic_content']")
	t.Replies, err = getReplyList(doc)
	if err != nil {
		return
	}
	return t, nil
}

func getReplyList(doc *html.Node) (s []Reply, err error) {
	for _, div := range htmlquery.Find(doc, "//*[@id='comment_list']/div") {
		r := Reply{}
		replyInfo := htmlquery.InnerText(htmlquery.FindOne(div, "./div[@class='re_info']"))
		r.Floor, r.Time, err = reInfo(replyInfo)
		if err != nil {
			return
		}
		r.Author = htmlquery.InnerText(htmlquery.FindOne(div, ".//div[@class='inner']/span//a/@href"))
		if htmlquery.FindOne(div, ".//div[@class='topic_sub_reply']") != nil {
			r.Replies, err = getSubReplyList(div)
			if err != nil {
				return
			}
		}
		s = append(s, r)
	}
	return
}

func getSubReplyList(doc *html.Node) (s []SubReply, err error) {
	for _, div := range htmlquery.Find(doc, "./div[@class='inner']//div[@class='topic_sub_reply']/div") {
		r := SubReply{}
		replyInfo := htmlquery.InnerText(htmlquery.FindOne(div, ".//div[@class='re_info']"))
		r.Floor, r.Time, err = reInfo(replyInfo)
		if err != nil {
			return
		}
		r.Author = htmlquery.InnerText(htmlquery.FindOne(div, ".//strong[@class='userName']//a/@href"))
		s = append(s, r)
	}
	return
}

func reInfo(s string) (floor Floor, t time.Time, err error) {
	// #1 - 2020-5-4 04:27
	// #1-1 - 2020-5-4 04:27
	s = strings.Trim(s, "\t\n #")
	l := strings.SplitN(s, " - ", 2)
	floorS := strings.Split(l[0], "-")
	floor[0], err = strconv.Atoi(floorS[0])
	if err != nil {
		return
	}
	if len(floorS) > 1 {
		floor[1], err = strconv.Atoi(floorS[1])
		if err != nil {
			return
		}
	}
	t, err = time.Parse("2006-1-2 15:04", l[1])
	return
}
