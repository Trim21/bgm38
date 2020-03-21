package main

import (
	"fmt"

	"bgm38/cmd"
	"bgm38/pkg/utils"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
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
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	cmd.Execute()
}
