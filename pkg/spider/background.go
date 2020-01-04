package spider

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"bgm38/config"
	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func dispatcher(urlToFetch chan string) {
	count := 0
	go func() {
		t := time.NewTicker(time.Duration(time.Second * 10))
		defer t.Stop()
		for {
			<-t.C
			fmt.Printf("[%s] dispatch %d items...\n", time.Now().Format("2006-01-02 15:04:05"), count)
		}
	}()

	logrus.Debugln("start spider dispatcher")
	for ; ; {
		res, err := db.Redis.BRPop(time.Minute, config.RedisSpiderURLKey).Result()
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		for _, s := range res {
			if s == config.RedisSpiderURLKey {
				continue
			}
			count++
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
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic occured: ", r)
				}
			}()

			subjectID, err := getSubjectID(res.url)
			if err != nil {
				logrus.Errorln(err)
				return
			}
			if subjectID == 0 {
				logrus.Fatalln(res.url)
			}
			doc, err := htmlquery.Parse(bytes.NewReader(res.res.Body()))
			if err != nil {
				logrus.Errorln(err)
				return
			}
			bodyString := string(res.res.Body())
			if strings.Contains("出错了", bodyString) {
				return
			}
			var subject db.Subject
			if strings.Contains("出错了", bodyString) {
				subject.Locked = 1
			}

			title := htmlquery.FindOne(doc, `//*[@id="headerSubject"]/h1/a`)
			if title == nil {
				return
			}
			subject.ID = subjectID
			subject.NameCn = htmlquery.SelectAttr(title, "title")
			subject.Name = htmlquery.InnerText(title)
			getTagFromDoc(doc, subjectID)
			getRelation(doc, subjectID)
			getEpList(doc, subjectID)
			getCollectorCount(doc, &subject)
			subject.Image = getImageURL(doc)
			subject.Score = getScore(doc)
			subject.Info = getInfo(doc)
			subject.SubjectType = getSubjectType(doc)
			uploadSubject(subject)
		}()
	}
}
