package spider

import (
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/pkg/db"
)

func parseEpList(doc *html.Node, subjectID int) {
	var eps []*db.Ep
	for _, ep := range htmlquery.Find(doc, `//*[@id="subject_detail"]//ul[@class="prg_list"]/li/a`) {
		href := htmlquery.SelectAttr(ep, "href")
		if href == "" {
			continue
		}
		sl := strings.Split(href, "/")
		epID, err := strconv.Atoi(sl[len(sl)-1])
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		eps = append(eps,
			&db.Ep{
				EpID:      epID,
				SubjectID: subjectID,
				Name:      htmlquery.SelectAttr(ep, "title"),
				Episode:   formatEp(ep),
			})
	}
	if len(eps) > 0 {
		uploadEps(eps)
	}
}

func formatEp(ep *html.Node) string {
	epText := htmlquery.InnerText(htmlquery.FindOne(ep, "./text()"))
	if e, err := strconv.Atoi(epText); err == nil {
		return strconv.Itoa(e)
	}
	return epText
}

func uploadEps(eps []*db.Ep) {
	for _, ep := range eps {
		_, err := epUpsertStmt.Exec(ep)
		if err != nil {
			logger.Error("upsert ep error", zap.String("err", err.Error()), zap.Object("ep", ep))
		}
	}
}
