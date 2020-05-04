package spider

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-redis/redis/v7"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"bgm38/config"
	"bgm38/pkg/db"
)

func dispatcher(urlToFetch chan string) {
	count := 0
	go func() {
		t := time.NewTicker(time.Minute * 1)
		defer t.Stop()
		for {
			<-t.C
			fmt.Printf("[%s] dispatch %d items...\n", time.Now().Format("2006-01-02 15:04:05"), count)
		}
	}()

	logger.Debug("start spider dispatcher")
	for i := 310000; i > 0; i-- {
		urlToFetch <- fmt.Sprintf("https://mirror.bgm.rin.cat/subject/%d", i)
		count++
	}

	for {
		res, err := db.Redis.BRPop(time.Minute, config.RedisSpiderURLKey).Result()
		if err != nil && err != redis.Nil {
			logger.Error(err.Error())
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
	logger.Info("start downloader")
	for url := range urlToFetch {
		logger.Debug(url)
		req := client.R()
		res, err := req.Get(url)
		if err != nil {

			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					go func() {
						logger.Debug("re-send request", zap.String("url", url))
						time.Sleep(5 * time.Second)
						urlToFetch <- url
					}()
					continue
				}
			}

			logger.Error("request error", zap.String("err", err.Error()))
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
					fmt.Println("panic occurred: ", r)
				}
			}()

			subjectID, err := getSubjectID(res.url)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			if subjectID == 0 {
				logger.Fatal(res.url)
			}
			doc, err := htmlquery.Parse(bytes.NewReader(res.res.Body()))
			if err != nil {
				logger.Error(err.Error())
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
			parseTagFromDoc(doc, subjectID)
			parseRelation(doc, subjectID)
			parseEpList(doc, subjectID)
			getCollectorCount(doc, &subject)
			subject.Image = getImageURL(doc)
			subject.Score = getScore(doc)
			subject.Info = getInfo(doc)
			subject.ScoreDetails = getScoreDetails(doc)
			subject.SubjectType = getSubjectType(doc)
			uploadSubject(&subject)
		}()
	}
}
