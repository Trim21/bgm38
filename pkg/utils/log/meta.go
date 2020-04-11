package log

import (
	"go.uber.org/zap/zapcore"
)

type Metadata struct {
	Beat    string
	Version string
}

func (s *Metadata) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("beat", s.Beat)
	enc.AddString("version", s.Version)
	return nil
}
