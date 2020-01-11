package cron

import (
	"bytes"
	"time"

	"bgm38/config"
	"bgm38/pkg/db"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func genWikiURL() {
	var urls = make(map[string]bool)

	res, err := client.R().Get("https://mirror.bgm.rin.cat/wiki")
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
				db.Redis.LPush(config.RedisSpiderURLKey, key)
			}
		}
	}
}
