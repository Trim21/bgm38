package config

import (
	"github.com/sirupsen/logrus"
	"bgm38/pkg/utils"
	"os"
	"path/filepath"
)

//var UpstreamIndex = "https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple/"
//var UpstreamIndex = "https://mirrors.aliyun.com/pypi/simple/"
var UpstreamIndex = "https://pypi.org/vote/"
var CacheDir string = utils.GetEnv("CACHE_DIR", "")
var RedisAddr = utils.GetEnv("REDIS_HOST", "127.0.0.1") + ":6379"
var RedisPassword = utils.GetEnv("REDIS_PASSWORD", "")

func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		logrus.Error(err)
		return
	}
	if CacheDir == "" {
		CacheDir = filepath.Join(home, ".cache", "gol")
	}

	_, err = os.Stat(CacheDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(CacheDir, os.ModeDir)
		if err != nil {
			logrus.Error(err)
		}
	}
}
