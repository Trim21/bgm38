package viewip

import (
	"encoding/json"

	"bgm38/web/app/db"
)

func formatData(subjects []db.Subject, relations []db.Relation) *subjectMap {
	edges := make([]edge, len(relations))
	nodes := make([]node, len(subjects))
	subjectIDMap := make(map[int]int)
	for index, subject := range subjects {
		var image = subject.Image
		if image == "" {
			image = "lain.bgm.tv/img/no_icon_subject.png"
		}
		image = "//" + image

		var info map[string][]string
		err := json.Unmarshal([]byte(subject.Info), &info)
		if err != nil {
			info = map[string][]string{}
		}
		subjectIDMap[subject.ID] = index
		nodes[index] = node{
			ID:          index,
			SubjectID:   subject.ID,
			Name:        subject.Name,
			NameCN:      subject.NameCn,
			Image:       image,
			SubjectType: subject.SubjectType,
			Info:        info,
			Map:         subject.Map,
		}
	}

	for index, relation := range relations {
		edges[index] = edge{
			ID:       relation.ID,
			Relation: relation.Relation,
			Source:   subjectIDMap[relation.Source],
			Target:   subjectIDMap[relation.Target],
			Map:      relation.Map,
		}
	}

	return &subjectMap{
		Edges: edges,
		Nodes: nodes,
	}
}
