package cron

import (
	"fmt"

	"bgm38/config"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type lo struct{}

func (l lo) Info(msg string, keysAndValues ...interface{}) {
	x := append([]interface{}{msg}, keysAndValues...)
	logrus.Infoln(x...)
}

func (l lo) Error(err error, msg string, keysAndValues ...interface{}) {
	x := append([]interface{}{err, msg}, keysAndValues...)
	logrus.Errorln(x...)
}

func Start() error {
	var err error
	fmt.Println("setup cron")
	var logger = lo{}
	c := cron.New(cron.WithLocation(config.TimeZone), cron.WithSeconds(), cron.WithLogger(logger), cron.WithChain(
		cron.Recover(logger), // or use cron.DefaultLogger
	))
	_, err = c.AddFunc("0 0 3 * * *", genWikiURL)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	_, err = c.AddFunc("0 0 3 1 * *", genWikiURL)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	_, err = c.AddFunc("0 0 3 3 * *", reCalculateMap)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	fmt.Println("start cron")
	c.Start()
	ch := make(chan bool)
	<-ch
	return nil
}
