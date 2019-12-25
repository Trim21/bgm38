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

func getRelation(doc *html.Node, subjectID int) {
	section := htmlquery.Find(doc, `//div[@class="subject_section"][//h2[@class="subtitle" and contains(text(), "关联条目")]]/div[@class="content_inner"]/ul/li`)
	var relations []db.Relation

	var chunkList = make([][]*html.Node, 0)

	for _, li := range section {
		if strings.Contains(htmlquery.SelectAttr(li, "class"), "sep") {
			chunkList = append(chunkList, []*html.Node{li})
		} else {
			chunkList[len(chunkList)-1] = append(chunkList[len(chunkList)-1], li)
		}

		for _, list := range chunkList {
			rel := htmlquery.InnerText(htmlquery.FindOne(list[0], "span/text()"))
			for _, li := range list {
				targetText := htmlquery.InnerText(htmlquery.FindOne(li, "a/@href"))
				sl := strings.Split(targetText, "/")
				target, err := strconv.Atoi(sl[len(sl)-1])
				if err != nil {
					logrus.Errorln(err)
					continue
				}
				relations = append(relations,
					db.Relation{
						ID:       fmt.Sprintf("%d-%d", subjectID, target),
						Relation: rel,
						Source:   subjectID,
						Target:   target,
					})
			}
		}
	}
	if len(relations) > 0 {
		uploadRelations(relations)
	}
}

func uploadRelations(relations []db.Relation) {
	var s []string
	var args []interface{}
	for _, r := range relations {
		s = append(s, "(?, ?, ?, ?)")
		args = append(args, r.ID, r.Relation, r.Source, r.Target)
	}
	sql := fmt.Sprintf("INSERT INTO `relation` (`id`, `relation`, `source`, `target`) VALUES %s ON DUPLICATE KEY UPDATE `relation` = VALUES(`relation`)", strings.Join(s, " , "))
	err := db.Mysql.Exec(sql, args...).Error
	if err != nil {
		logrus.Errorln(err, sql, args)
	}

}
