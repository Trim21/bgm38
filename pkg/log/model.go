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
	MetaData   MetaData `msg:"@metadata"`
	Msg        string   `msg:"msg"`
	LoggedAt   int64    `msg:"logged_at"`
	LineNumber int      `msg:"line_number"`
	Function   string   `msg:"function"`
	Level      string   `msg:"level"`
	Module     string   `msg:"module"`
	Process    int      `msg:"process"`
}

func fromEntry(entry *logrus.Entry) (*LoggedMessage, error) {
	var msg = &LoggedMessage{
		MetaData: meta,
		Msg:      entry.Message,
		LoggedAt: entry.Time.Unix(),
		Level:    entry.Level.String(),
		Process:  pid,
	}
	if entry.HasCaller() {
		msg.LineNumber = entry.Caller.Line
		msg.Function = entry.Caller.Function
		msg.Module = entry.Caller.Function
	}
	return msg, nil
}
