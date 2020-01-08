package main

import (
	"fmt"
	"path"
	"runtime"

	"bgm38/cmd"
	"bgm38/pkg/cron"
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

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	cron.Init()

	cmd.Execute()
}
