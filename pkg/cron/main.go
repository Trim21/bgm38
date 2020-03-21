package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"bgm38/config"
	"bgm38/pkg/db"
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

func Run(name string) error {
	db.InitDB()
	switch name {
	case "reCalculateMap":
		reCalculateMap()
	case "genWikiURL":
		genWikiURL()
	case "genFullURL":
		genFullURL()
	default:
		return fmt.Errorf("cron name not exist")
	}
	return nil
}

func Start() error {
	db.InitDB()
	var err error
	fmt.Println("setup cron")
	var logger = lo{}
	c := cron.New(cron.WithLocation(config.TimeZone), cron.WithSeconds(), cron.WithLogger(logger), cron.WithChain(
		cron.Recover(logger), // or use cron.DefaultLogger
	))
	_, err = c.AddFunc("0 0 3 * * *", genFullURL)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	_, err = c.AddFunc("0 0 3 1 * *", genWikiURL)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	_, err = c.AddFunc("0 0 3 2 * *", reCalculateMap)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	fmt.Println("start cron")
	c.Run()
	return nil
}
