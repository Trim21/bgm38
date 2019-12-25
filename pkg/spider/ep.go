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

func getEpList(doc *html.Node, subjectID int) {
	var eps []db.Ep
	for _, ep := range htmlquery.Find(doc, `//*[@id="subject_detail"]//ul[@class="prg_list"]/li/a`) {
		href := htmlquery.SelectAttr(ep, "href")
		if href == "" {
			continue
		}
		sl := strings.Split(href, "/")
		epID, err := strconv.Atoi(sl[len(sl)-1])
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		eps = append(eps,
			db.Ep{
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

func uploadEps(eps []db.Ep) {
	var s []string
	var args []interface{}
	for _, ep := range eps {
		s = append(s, "(?, ?, ?, ?)")
		args = append(args, ep.EpID, ep.SubjectID, ep.Name, ep.Episode)
	}
	sql := fmt.Sprintf("INSERT INTO `ep` (`ep_id`, `subject_id`, `name`, `episode`) VALUES %s ON DUPLICATE KEY UPDATE `name` = VALUES(`name`), `episode` = VALUES(`episode`)", strings.Join(s, ","))
	err := db.Mysql.Exec(sql, args...).Error
	if err != nil {
		logrus.Errorln(err, sql, args)
	}
}
