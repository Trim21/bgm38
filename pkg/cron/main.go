package cron

import (
	"fmt"
	"strconv"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"bgm38/config"
	"bgm38/pkg/db"
	"bgm38/pkg/utils/log"
)

var logger *zap.Logger

type lo struct{}

func formatKeysAndValues(keysAndValues ...interface{}) []zap.Field {
	var s []zap.Field
	for i, value := range keysAndValues {
		switch v := value.(type) {
		case int:
			s = append(s, zap.Int(strconv.Itoa(i), v))
		case string:
			s = append(s, zap.String(strconv.Itoa(i), v))
		case bool:
			s = append(s, zap.Bool(strconv.Itoa(i), v))
		default:
			s = append(s, zap.String(strconv.Itoa(i), fmt.Sprintf("%v", v)))
		}
	}
	return s
}

func (l lo) Info(msg string, keysAndValues ...interface{}) {
	logger.Info(msg, formatKeysAndValues(keysAndValues)...)
}

func (l lo) Error(err error, msg string, keysAndValues ...interface{}) {
	s := []zap.Field{zap.String("err", err.Error())}
	s = append(s, formatKeysAndValues(keysAndValues)...)
	logger.Error(msg, s...)
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
	logger = log.BindMeta("bgm38-cron-v1", log.CreateLogger())
	var err error
	fmt.Println("setup cron")
	var logger = lo{}
	c := cron.New(cron.WithLocation(config.TimeZone), cron.WithSeconds(), cron.WithLogger(logger), cron.WithChain(
		cron.Recover(logger), // or use cron.DefaultLogger
	))

	_, err = c.AddFunc("0 0 3 * * *", genFullURL)
	checkErr(err)

	_, err = c.AddFunc("0 0 3 1 * *", genWikiURL)
	checkErr(err)

	_, err = c.AddFunc("0 0 3 2 * *", reCalculateMap)
	checkErr(err)

	fmt.Println("start cron")
	c.Run()
	return nil
}
func checkErr(err error) {
	if err != nil {
		logger.Fatal(err.Error())
	}
}
