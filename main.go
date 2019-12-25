package main

import (
	"fmt"
	"path"
	"runtime"

	"bgm38/cmd"
	"bgm38/pkg/utils"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

func main() {
	if utils.GetEnv("DEBUG", "0") == "1" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})

	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "___DSN___",
	})

	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	cmd.Execute()
}
