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
	Usergroup    int    `json:"usergroup"`
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
	SubjectID int    `gorm:"primary_key" json:"-"`
	Text      string `gorm:"type:char;primary_key"`
	Count     int    `json:"count"`
}

//
type Subject struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	Image        string `json:"image" gorm:"type:varchar(255)"`
	SubjectType  string `json:"subject_type" gorm:"type:varchar(255)"`
	NameCn       string `json:"name_cn" gorm:"type:varchar(255)"`
	Tags         string `json:"tags"`
	Info         string `json:"info"`
	ScoreDetails string `json:"score_details"`
	Score        string `json:"score" gorm:"type:varchar(255)"`
	Wishes       int    `json:"wishes"`
	Done         int    `json:"done"`
	Doings       int    `json:"doings"`
	OnHold       int    `json:"on_hold"`
	Dropped      int    `json:"dropped"`
	Map          int    `json:"-"`
	Locked       int8   `json:"-"`
}

//

type Relation struct {
	ID       string `gorm:"type:varchar(255);primary_key" json:"-"`
	Relation string `json:"relation" gorm:"type:varchar(255)"`
	Source   int    `json:"source"`
	Target   int    `json:"target"`
	Map      int    `json:"-"`
	Removed  int8   `json:"-"`
}

//
type MissingBangumi struct {
	Source    string `gorm:"type:char;primary_key" json:"-"`
	BangumiID string `gorm:"primary_key;'bangumi_id'" json:"-"`
}
