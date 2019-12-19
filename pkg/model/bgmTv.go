package model

type SubjectCollection struct {
	Wish    int `json:"wish"`
	Collect int `json:"collect"`
	Doing   int `json:"doing"`
	OnHold  int `json:"on_hold"`
	Dropped int `json:"dropped"`
}

type SubjectImage struct {
	Large  string `json:"large"`
	Common string `json:"common"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
	Grid   string `json:"grid"`
}

type UserCollectionSubject struct {
	Images     SubjectImage      `json:"images"`
	ID         int               `json:"id"`
	URL        string            `json:"url"`
	Type       int               `json:"type"`
	Summary    string            `json:"summary"`
	Name       string            `json:"name"`
	NameCn     string            `json:"name_cn"`
	AirWeekday int               `json:"air_weekday"`
	AirDate    string            `json:"air_date"`
	Eps        int               `json:"eps"`
	EpsCount   int               `json:"eps_count"`
	Collection SubjectCollection `json:"collection"`
}

type UserCollection struct {
	Name      string                `json:"name"`
	SubjectID int                   `json:"subject_id"`
	EpStatus  int                   `json:"ep_status"`
	VolStatus int                   `json:"vol_status"`
	LastTouch int                   `json:"lasttouch"`
	Subject   UserCollectionSubject `json:"subject"`
}
