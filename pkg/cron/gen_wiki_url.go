package cron

import (
	"bytes"
	"fmt"
	"net/url"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"

	"bgm38/config"
	"bgm38/pkg/db"
)

var client = resty.New()

func genWikiURL() {
	var urls = make(map[string]bool)
	pageUrl, _ := url.Parse("https://mirror.bgm.rin.cat/wiki")
	res, err := client.R().Get(pageUrl.String())
	if err != nil {
		time.Sleep(5 * time.Second)
		genWikiURL()
		return
	}
	doc, err := htmlquery.Parse(bytes.NewReader(res.Body()))
	if err != nil {
		time.Sleep(5 * time.Second)
		genWikiURL()
		return
	}

	for _, url := range htmlquery.Find(doc, `//*[@id="wikiEntryMainTab"]//li/a/@href`) {
		urls[htmlquery.InnerText(url)] = true
	}
	for _, url := range htmlquery.Find(doc, `//*[@id="latestEntryMainTab"]//li/a/@href`) {
		urls[htmlquery.InnerText(url)] = true

	}

	for key, value := range urls {
		if value {
			if key != "" {
				u, _ := url.Parse(key)
				fmt.Println(pageUrl.ResolveReference(u).String())
				db.Redis.LPush(config.RedisSpiderURLKey, pageUrl.ResolveReference(u).String())
			}
		}
	}
}
