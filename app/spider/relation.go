package spider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/pkg/db"
)

func parseRelation(doc *html.Node, subjectID int) {
	section := htmlquery.Find(doc, `//div[@class="subject_section"]`+
		`[//h2[@class="subtitle" and contains(text(), "关联条目")]]`+
		`/div[@class="content_inner"]/ul/li`)
	var relations []*db.Relation

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
					logger.Error(err.Error())
					continue
				}
				relations = append(relations,
					&db.Relation{
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

func uploadRelations(relations []*db.Relation) {
	var err error
	for _, relation := range relations {
		_, err = relationUpsertStmt.Exec(relation)
		if err != nil {
			logger.Error("upsert into table `relation`",
				zap.Object("relation", relation),
				zap.String("err", err.Error()))
		}
	}

}
