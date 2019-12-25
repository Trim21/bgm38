package spider

import (
	"bytes"
	"strings"
	"time"

	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func dispatcher(urlToFetch chan string) {
	logrus.Debugln("start spider dispatcher")
	for ; ; {
		res, err := db.Redis.BRPop(time.Minute, redisKey).Result()
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		for _, s := range res {
			if s == redisKey {
				continue
			}
			urlToFetch <- s
		}
	}
}

func downloader(urlToFetch chan string, resQueue chan response) {
	var client = resty.New()
	logrus.Info("start downloader")
	for url := range urlToFetch {
		logrus.Debugln(url)
		req := client.R()
		res, err := req.Get(url)
		if err != nil {
			urlToFetch <- url
			continue
		}
		if res.StatusCode() > 300 || bytes.Contains([]byte("502 Bad Gateway"), res.Body()) {
			go func() {
				time.Sleep(5 * time.Second)
				urlToFetch <- url
			}()
		}
		resQueue <- response{
			url: url,
			res: res,
		}

	}
}

func parser(resQueue chan response) {
	for res := range resQueue {

		subjectID, err := getSubjectID(res.url)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		if subjectID == 0 {
			logrus.Fatalln(res.url)
		}
		doc, err := htmlquery.Parse(bytes.NewReader(res.res.Body()))
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		bodyString := string(res.res.Body())
		if strings.Contains("出错了", bodyString) {
			continue
		}
		var subject db.Subject
		if strings.Contains("出错了", bodyString) {
			subject.Locked = 1
		}

		title := htmlquery.FindOne(doc, `//*[@id="headerSubject"]/h1/a`)
		if title == nil {
			continue
		}
		subject.ID = subjectID
		subject.NameCn = htmlquery.SelectAttr(title, "title")
		subject.Name = htmlquery.InnerText(title)
		getTagFromDoc(doc, subjectID)
		getRelation(doc, subjectID)
		getEpList(doc, subjectID)
		getCollectorCount(doc, &subject)
		subject.Image = getImageUrl(doc)
		subject.Score = getScore(doc)
		subject.Info = getInfo(doc)
		subject.SubjectType = getSubjectType(doc)
		uploadSubject(subject)
	}
}
