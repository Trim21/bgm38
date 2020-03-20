package db

import "time"

//
type Usertoken struct {
	UserID       int    `gorm:"primary_key;'user_id'" json:"-"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AuthTime     int    `json:"auth_time"`
	AccessToken  string `json:"access_token" gorm:"type:char"`
	RefreshToken string `json:"refresh_token" gorm:"type:char"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	UserGroup    int    `json:"usergroup"`
}

//
type UserSubmitBangumi struct {
	Source     string    `gorm:"type:char;primary_key" json:"-"`
	SubjectID  int       `json:"subject_id" gorm:"'subject_id'"`
	BangumiID  string    `gorm:"primary_key;'bangumi_id'" json:"-"`
	UserID     int       `gorm:"primary_key;'user_id'" json:"-"`
	ModifyTime time.Time `json:"modify_time"`
}

//
type Tag struct {
	SubjectID int    `gorm:"primary_key" json:"-" db:"subject_id"`
	Text      string `gorm:"type:char;primary_key" db:"text"`
	Count     int    `json:"count" db:"count"`
}

//
type Subject struct {
	ID           int    `gorm:"primary_key" json:"id" db:"id"`
	Name         string `json:"name" gorm:"type:varchar(255)" db:"name"`
	Image        string `json:"image" gorm:"type:varchar(255)" db:"image"`
	SubjectType  string `json:"subject_type" gorm:"type:varchar(255)" db:"subject_type"`
	NameCn       string `json:"name_cn" gorm:"type:varchar(255)" db:"name_cn"`
	Tags         string `json:"tags" db:"tags"`
	Info         string `json:"info" db:"info"`
	ScoreDetails string `json:"score_details" db:"score_details"`
	Score        string `json:"score" gorm:"type:varchar(255)" db:"score"`
	Wishes       int    `json:"wishes" db:"wishes"`
	Done         int    `json:"done" db:"done"`
	Doings       int    `json:"doings" db:"doings"`
	OnHold       int    `json:"on_hold" db:"on_hold"`
	Dropped      int    `json:"dropped" db:"dropped"`
	Map          int    `json:"-" db:"map"`
	Locked       int8   `json:"-" db:"locked"`
}

//

type Relation struct {
	ID       string `gorm:"type:varchar(255);primary_key" json:"-" db:"id"`
	Relation string `json:"relation" gorm:"type:varchar(255)" db:"relation"`
	Source   int    `json:"source" db:"source"`
	Target   int    `json:"target" db:"target"`
	Map      int    `json:"-" db:"map"`
	Removed  int8   `json:"-" db:"removed"`
}

//
type MissingBangumi struct {
	Source    string `gorm:"type:char;primary_key" json:"-"`
	BangumiID string `gorm:"primary_key;'bangumi_id'" json:"-"`
}
