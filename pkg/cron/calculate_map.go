package cron

import (
	"fmt"

	"bgm38/pkg/db"
)

const (
	subjectTable  = "subject"
	relationTable = "relation"
	typeMusic     = "Music"
	chunkSize     = 5000
)

var maxMapID = 0
var blankList = []string{"角色出演", "片头曲", "片尾曲", "其他", "画集", "原声集"}

func removeRelation(source int, target int) {
	db.Mysql.Updates(db.Relation{Removed: 1}).Where(`id = ? or id = ?`, fmt.Sprintf("%d-%d", source, target), fmt.Sprintf("%d-%d", target, source))
}

func reCalculateMap() {

}

func removeNodes(nodeIDs ...int) {
	db.Mysql.Model(&db.Subject{}).UpdateColumn(`locked`, 1).Where(`id in ?`, nodeIDs)
	db.Mysql.Model(&db.Relation{}).UpdateColumn(`removed`, 1).Where(`target in ? or source in `, nodeIDs, nodeIDs)
}

func relationsNeedToRemove(m map[int]int) {
	for id1, id2 := range m {
		db.Mysql.Model(&db.Relation{}).UpdateColumn(`removed`, 1).
			Where(`target = ? or target = ? or source = ? or source = ?`, id1, id2, id1, id2)
	}
}

func preRemoveRelation() {
	db.Mysql.Model(&db.Relation{}).UpdateColumn(`removed`, 1).
		Where(`relation in `, blankList)
}

func preRemove(subjectStart int, subjectEnd int) {

	println("pre remove")
	removeNodes(91493, 102098, 228714, 231982, 932, 84944, 78546)
	preRemoveRelation()
	relationsNeedToRemove(map[int]int{
		91493:  8,
		8108:   35866,
		446:    123207,
		123207: 27466,
		123217: 4294,
	})

	var idToRemove = make(map[int]bool)
	db.Mysql.Table(subjectTable).Update(`locked`, 1).Where(`subject_type = ?`, typeMusic)
	var subjects []db.Subject
	db.Mysql.Where(`locked = ?`, 1).Find(&subjects)
	for _, s := range subjects {
		idToRemove[s.ID] = true

	}

	var idToRM []int
	for key := range idToRemove {
		idToRM = append(idToRM, key)
	}

	db.Mysql.Table(relationTable).Update(`removed`, 1).Where(`source in ? or target in ?`, idToRM, idToRM)

	for i := subjectStart; i < subjectEnd; i += chunkSize {
		relationIDNeedToRemove := make(map[string]bool)
		sourceToTarget := make(map[int]map[int]bool)
		var rels [] db.Relation
		db.Mysql.Where(`(( source >= ? AND source < ? ) OR ( target >= ? AND target < ? ) ) AND removed = ?`,
			i, i+chunkSize, i, i+chunkSize, 0).Find(&rels)

		for _, rel := range rels {
			sourceToTarget[rel.Source][rel.Target] = true
		}

		for _, rel := range rels {
			if subMap, ok := sourceToTarget[rel.Target]; ok {
				if _, ok := subMap[rel.Source]; ok {
					continue
				}
			}
			relationIDNeedToRemove[rel.ID] = true
		}
	}
}

func chunkIter(s []int, chunkSize int, f func([]int)) {
	l := len(s)
	for i := 0; i < l; i = i + chunkSize {
		f(s[i:min(i+chunkSize, l)])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pyRange(start, end, step int) []int {
	s := make([]int, 0, (end-start)/step+3)
	for i := start; i < end; i += step {
		s = append(s, i)
	}
	return s
}
