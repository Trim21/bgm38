package spider

import (
	"strconv"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/pkg/db"
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
			logger.Error(err.Error())
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
			logger.Error("insert into table `tag`",
				zap.Object("tag", tag), zap.String("err", err.Error()))
		}
	}
}
