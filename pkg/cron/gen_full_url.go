package cron

import (
	"fmt"

	"bgm38/config"
	"bgm38/pkg/db"
)

func genFullURL() {
	var urls []interface{}
	for i := 1; i < 310000; i++ {
		urls = append(urls, fmt.Sprintf("https://mirror.bgm.rin.cat/subject/%d", i))
		if i%500 == 0 {
			db.Redis.LPush(config.RedisSpiderURLKey, urls...)
			urls = nil
		}
	}
	db.Redis.LPush(config.RedisSpiderURLKey, urls...)
}

