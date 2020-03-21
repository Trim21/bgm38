package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"

	"bgm38/cmd"
	"bgm38/config"
	"bgm38/pkg/log"
	"bgm38/pkg/utils"
)

func main() {
	if utils.GetEnv("DEBUG", "0") != "0" {
		fmt.Println("set logger to debug level")
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "___DSN___",
		})

		if err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}
		hook := log.NewRedisHook(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			PoolSize: 3,
		}, "bgm38 log")
		logrus.AddHook(hook)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	cmd.Execute()
}
