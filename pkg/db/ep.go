package db

import (
	"time"

	"go.uber.org/zap/zapcore"
)

//
type Ep struct {
	EpID      int        `gorm:"primary_key" json:"-" db:"ep_id"`
	SubjectID int        `json:"subject_id" db:"subject_id"`
	Name      string     `json:"name" db:"name"`
	Episode   string     `json:"episode" db:"episode"`
	AirTime   *time.Time `json:"air_time" db:"air_time"`
}

func (e *Ep) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("id", e.EpID)
	enc.AddInt("subject_id", e.SubjectID)
	enc.AddString("name", e.Name)
	enc.AddString("episode", e.Episode)
	if e.AirTime != nil {
		enc.AddTime("air_time", *e.AirTime)
	} else {
		enc.AddString("air_time", "null")
	}
	return nil
}

//
type EpBilibili struct {
	SourceEpID int    `gorm:"primary_key" json:"-"`
	EpID       int    `json:"ep_id"`
	SubjectID  int    `json:"subject_id"`
	Title      string `json:"title"`
}

//
type EpIqiyi struct {
	SourceEpID string `gorm:"primary_key" json:"-"`
	EpID       int    `json:"ep_id"`
	SubjectID  int    `json:"subject_id"`
	Title      string `json:"title"`
}

//
type EpSource struct {
	SubjectID  int    `json:"subject_id"`
	Source     string `gorm:"type:char;primary_key" json:"-"`
	SourceEpID string `gorm:"primary_key" json:"-"`
	BgmEpID    int    `json:"bgm_ep_id"`
	Episode    int    `json:"episode"`
}
