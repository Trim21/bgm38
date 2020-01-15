package spider

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
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
		el := htmlquery.FindOne(doc, fmt.Sprintf(`//*[@id="subjectPanelCollect"]/span[@class="tip_i"]/a[@href="/subject/%d/%s"]/text()`, subject.ID, value))
		if el == nil {
			continue
		}
		s := htmlquery.InnerText(el)

		v, err := strconv.Atoi(strings.Split(s, "人")[0])
		if err != nil {
			logrus.Errorln(err)
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

func getInfo(doc *html.Node) string {
	var info = make(map[string][]string)

	for _, el := range htmlquery.Find(doc, `//*[@id="infobox"]/li`) {
		key := strings.Replace(htmlquery.InnerText(htmlquery.FindOne(el, "span/text()")), ": ", "", -1)
		var value []string
		for _, node := range htmlquery.Find(el, "a/text()") {
			value = append(value, htmlquery.InnerText(node))
		}

		for _, node := range htmlquery.Find(el, "text()") {
			text := htmlquery.InnerText(node)
			if text == "、" {
				continue
			} else if strings.HasSuffix(text, "、") {
				value = append(value, strings.TrimRight(text, "、"))
			} else {
				value = append(value, text)
			}
		}

		info[key] = value

	}
	s, _ := json.Marshal(info)
	return string(s)
}

func getSubjectType(doc *html.Node) string {
	subjectType := htmlquery.InnerText(htmlquery.FindOne(doc, `//*[@id="panelInterestWrapper"]//div[contains(@class,"global_score")]/div/small[contains(@class, "grey")]/text()`))
	sl := strings.Split(subjectType, " ")
	return sl[1]

}

func uploadSubject(subject db.Subject) {
	op := "ON DUPLICATE KEY UPDATE `name_cn` = VALUES(`name_cn`), `name` = VALUES(`name`), `image` = VALUES(`image`), `tags` = VALUES(`tags`), `locked` = VALUES(`locked`), `info` = VALUES(`info`), `score_details` = VALUES(`score_details`), `score` = VALUES(`score`), `wishes` = VALUES(`wishes`), `done` = VALUES(`done`), `doings` = VALUES(`doings`), `on_hold` = VALUES(`on_hold`), `dropped` = VALUES(`dropped`)"
	err := db.Mysql.Set("gorm:insert_option", op).Create(&subject).Error
	if err != nil {
		logrus.Errorln(err)
	}
}
