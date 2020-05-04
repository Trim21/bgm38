package spider

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"

	"bgm38/pkg/db"
)

var collector = map[string]string{
	"wishes":  "wishes",
	"done":    "collections",
	"doings":  "doings",
	"on_hold": "on_hold",
	"dropped": "dropped",
}

func getCollectorCount(doc *html.Node, subject *db.Subject) {
	for key, value := range collector {
		el := htmlquery.FindOne(doc, fmt.Sprintf(`//*[@id="subjectPanelCollect"]`+
			`/span[@class="tip_i"]/a[@href="/subject/%d/%s"]/text()`, subject.ID, value))
		if el == nil {
			continue
		}
		s := htmlquery.InnerText(el)

		v, err := strconv.Atoi(strings.Split(s, "人")[0])
		if err != nil {
			logger.Error(err.Error())
		}

		switch key {
		case "wishes":
			subject.Wishes = v
		case "done":
			subject.Done = v
		case "doings":
			subject.Doings = v
		case "on_hold":
			subject.OnHold = v
		case "dropped":
			subject.Dropped = v
		}
	}
}

func getScoreDetails(doc *html.Node) string {
	var detail = make(map[string]string)
	detail["total"] = htmlquery.InnerText(htmlquery.FindOne(doc, `//*[@id="ChartWarpper"]/div/small/span/text()`))
	for _, li := range htmlquery.Find(doc, `//*[@id="ChartWarpper"]/ul[@class="horizontalChart"]/li`) {
		s := htmlquery.InnerText(htmlquery.FindOne(doc, `.//span[@class="count"]/text()`))
		detail[htmlquery.InnerText(htmlquery.FindOne(li, `.//span[@class="label"]/text()`))] = s[1 : len(s)-1]
	}
	data, err := json.Marshal(detail)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func getInfo(doc *html.Node) string {
	var info = make(map[string][]string)

	for _, el := range htmlquery.Find(doc, `//*[@id="infobox"]/li`) {
		key := strings.Replace(htmlquery.InnerText(htmlquery.FindOne(el, "span/text()")), ": ", "", -1)
		var value []string
		for _, node := range htmlquery.Find(el, "a/text()") {
			value = append(value, htmlquery.InnerText(node))
		}

		for _, node := range htmlquery.Find(el, "text()") {
			switch text := htmlquery.InnerText(node); {
			case text == "、":
				continue
			case strings.HasSuffix(text, "、"):
				value = append(value, strings.TrimRight(text, "、"))
			default:
				value = append(value, text)
			}
		}

		info[key] = value

	}
	s, _ := json.Marshal(info)
	return string(s)
}

func getSubjectType(doc *html.Node) string {
	subjectType := htmlquery.InnerText(htmlquery.FindOne(doc,
		`//*[@id="panelInterestWrapper"]//div[contains(@class,"global_score")]`+
			`/div/small[contains(@class, "grey")]/text()`))
	sl := strings.Split(subjectType, " ")
	return sl[1]

}

func uploadSubject(subject *db.Subject) {
	_, err := subjectUpsertStmt.Exec(subject)
	if err != nil {
		logrus.Errorln(err)
	}
}
