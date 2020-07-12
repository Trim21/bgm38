package spider

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/config"
	"bgm38/pkg/db"
)

var airTimePattern = regexp.MustCompile(`首播:(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)

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
		moreInfoElSelector := htmlquery.SelectAttr(ep, "rel")
		moreInfoEl := htmlquery.FindOne(doc, fmt.Sprintf(`//*[@id="%s"]`, moreInfoElSelector[1:]))
		tip := htmlquery.InnerText(htmlquery.FindOne(moreInfoEl, `//span[@class="tip"]`))
		var airTime *time.Time = nil
		result := airTimePattern.FindStringSubmatch(tip)
		if len(result) > 0 {
			var e error
			year, err := strconv.Atoi(result[1])
			if err != nil {
				e = err
			}
			month, err := strconv.Atoi(result[2])
			if err != nil {
				e = err
			}
			day, err := strconv.Atoi(result[3])
			if err != nil {
				e = err
			}
			if e == nil {
				d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, config.TimeZone)
				airTime = &d
			}

		}
		eps = append(eps,
			&db.Ep{
				EpID:      epID,
				SubjectID: subjectID,
				Name:      htmlquery.SelectAttr(ep, "title"),
				Episode:   formatEp(ep),
				AirTime:   airTime,
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
