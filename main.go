package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"

	"bgm38/cmd"
	"bgm38/config"
	"bgm38/pkg/utils"
)

func main() {
	if utils.GetEnv("DEBUG", "0") == "0" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: config.DSN,
		})

		if err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}
	}
	cmd.Execute()
}
