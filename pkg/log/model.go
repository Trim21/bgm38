package log

import (
	"os"

	"bgm38/config"
	"github.com/sirupsen/logrus"
)

var pid = os.Getpid()
var meta = MetaData{
	Beat:    "app_logging",
	Version: config.Version,
}

type MetaData struct {
	Beat    string `msg:"beat"`
	Version string `msg:"version"`
}
type LoggedMessage struct {
	MetaData MetaData               `msg:"@metadata"`
	Msg      string                 `msg:"msg"`
	LoggedAt int64                  `msg:"logged_at"`
	Level    string                 `msg:"level"`
	Process  int                    `msg:"process"`
	Data     map[string]interface{} `msg:"data"`
}

func fromEntry(entry *logrus.Entry) *LoggedMessage {
	var msg = &LoggedMessage{
		MetaData: meta,
		Msg:      entry.Message,
		LoggedAt: entry.Time.Unix(),
		Level:    entry.Level.String(),
		Process:  pid,
		Data:     entry.Data,
	}
	return msg
}
