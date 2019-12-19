package viewip

type edge struct {
	ID       string `json:"id"`
	Relation string `json:"relation"`
	Source   int    `json:"source"`
	Target   int    `json:"target"`
	Map      int    `json:"map"`
}

type node struct {
	ID          int                 `json:"id"`
	SubjectID   int                 `json:"subject_id"`
	Name        string              `json:"name"`
	NameCN      string              `json:"name_cn"`
	Image       string              `json:"image"` // = 'lain.bgm.tv/img/no_icon_subject.png'
	SubjectType string              `json:"subject_type"`
	Info        map[string][]string `json:"info"`
	Map         int                 `json:"map"`
}

type subjectMap struct {
	Nodes []node `json:"nodes"`
	Edges []edge `json:"edges"`
}
