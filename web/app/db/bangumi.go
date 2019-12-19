package db

//
type BangumiBilibili struct {
	SubjectID int    `gorm:"primary_key" json:"-"`
	MediaID   int    `json:"media_id"`
	SeasonID  int    `json:"season_id"`
	Title     string `json:"title"`
}

//
type BangumiIqiyi struct {
	SubjectID int    `gorm:"primary_key" json:"-"`
	BangumiID string `json:"bangumi_id"`
	Title     string `json:"title"`
}

//
type BangumiSource struct {
	Source    string `gorm:"type:char;primary_key" json:"-"`
	BangumiID string `gorm:"primary_key" json:"-"`
	SubjectID int    `json:"subject_id"`
}
