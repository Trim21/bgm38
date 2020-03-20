package spider

import (
	"strconv"

	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func parseTagFromDoc(doc *html.Node, subjectID int) {
	tags := htmlquery.Find(doc, `//*[@id="subject_detail"]//div[@class="subject_tag_section"]/div[@class="inner"]/a`)
	tgs := make([]*db.Tag, 0, len(tags))
	for _, n := range tags {
		el := htmlquery.FindOne(n, "span/text()")
		if el == nil {
			continue
		}
		text := htmlquery.InnerText(el)
		countText := htmlquery.InnerText(htmlquery.FindOne(n, "small/text()"))
		count, err := strconv.Atoi(countText)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		if text != "" {
			tgs = append(tgs, &db.Tag{
				SubjectID: subjectID,
				Text:      text,
				Count:     count,
			})
		}
	}
	if len(tgs) != 0 {
		uploadTags(tgs)
	}
}


func uploadTags(tags []*db.Tag) {
	var err error

	for _, tag := range tags {
		_, err = tagUpsertStmt.Exec(tag)
		if err != nil {
			logrus.Errorln("insert into table `tag`", tag, err)
		}
	}
}
