package spider

import (
	"fmt"
	"strconv"
	"strings"

	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func getTagFromDoc(doc *html.Node, subjectID int) {
	tags := htmlquery.Find(doc, `//*[@id="subject_detail"]//div[@class="subject_tag_section"]/div[@class="inner"]/a`)
	tgs := make([]db.Tag, len(tags))
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
			tgs = append(tgs, db.Tag{
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

func uploadTags(tags []db.Tag) {
	var s []string
	var args []interface{}
	for _, tag := range tags {
		s = append(s, "(?, ?, ?)")
		args = append(args, tag.SubjectID, tag.Text, tag.Count)
	}
	sql := fmt.Sprintf("INSERT INTO `tag` (`subject_id`, `text`, `count`) VALUES %s ON DUPLICATE KEY UPDATE `count` = VALUES(`count`)", strings.Join(s, " , "))
	err := db.Mysql.Exec(sql, args...).Error
	if err != nil {
		logrus.Errorln(err, sql, args)

	}

}
