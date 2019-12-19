package db

//
type Ep struct {
	EpID      int    `gorm:"primary_key" json:"-"`
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
	Episode   string `json:"episode"`
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
