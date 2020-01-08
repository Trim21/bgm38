package cron

import (
	"fmt"
	"sync"

	"bgm38/pkg/db"
	"github.com/sirupsen/logrus"
)

const (
	subjectTable  = "subject"
	relationTable = "relation"
	typeMusic     = "Music"
	chunkSize     = 5000
)

var m = sync.Mutex{}
var maxMapID = 0
var blankList = []string{"角色出演", "片头曲", "片尾曲", "其他", "画集", "原声集"}

func removeRelation(source int, target int) {
	db.Mysql.Where(`id = ? or id = ?`,
		fmt.Sprintf("%d-%d", source, target),
		fmt.Sprintf("%d-%d", target, source)).Updates(db.Relation{Removed: 1})
}

func reCalculateMap() {
	db.InitDB()
	var maxSubject = db.Subject{}
	var minSubject = db.Subject{}

	db.Mysql.Order(`id desc`).First(&maxSubject)
	db.Mysql.Order(`id`).First(&minSubject)
	db.Mysql.Table(subjectTable).Where("1 = ?", 1).Update(`locked`, 0)
	db.Mysql.Table(relationTable).Where("1 = ?", 1).Update(`removed`, 0)
	preRemove(minSubject.ID, maxSubject.ID)
	firstRun(minSubject.ID, maxSubject.ID)
}

func removeNodes(nodeIDs ...int) {
	db.Mysql.Table(subjectTable).Where(`id in (?)`, nodeIDs).Update(`locked`, 1)
	db.Mysql.Table(relationTable).Where(`target in (?) or source in (?)`, nodeIDs, nodeIDs).Update(`removed`, 1)
}

func relationsNeedToRemove(m map[int]int) {
	for id1, id2 := range m {
		db.Mysql.Model(&db.Relation{}).
			Where(`target = ? or target = ? or source = ? or source = ?`, id1, id2, id1, id2).
			Update(`removed`, 1)

	}
}

func preRemove(subjectStart int, subjectEnd int) {

	println("pre remove")
	removeNodes(91493, 102098, 228714, 231982, 932, 84944, 78546)

	db.Mysql.Table(relationTable).
		Where(`relation in (?)`, blankList).Update(`removed`, 1)

	relationsNeedToRemove(map[int]int{
		91493:  8,
		8108:   35866,
		446:    123207,
		123207: 27466,
		123217: 4294,
	})
	db.Mysql.Table(subjectTable).Where(`subject_type = ?`, typeMusic).Update(`locked`, 1)
	var idToRemove = make(map[int]bool)
	var subjects []db.Subject
	db.Mysql.Where(`locked = ?`, 1).Find(&subjects)
	for _, s := range subjects {
		idToRemove[s.ID] = true

	}

	var idToRM []int
	for key := range idToRemove {
		idToRM = append(idToRM, key)
	}

	err := chunkIterInt(idToRM, func(s []int) error {
		return db.Mysql.Table(relationTable).
			Where(`source IN (?) OR target IN (?)`, s, s).
			Update(`removed`, 1).Error
	})

	if err != nil {
		logrus.Errorln(err)
	}

	for i := subjectStart; i < subjectEnd; i += chunkSize {
		relationIDNeedToRemove := make(map[string]bool)
		sourceToTarget := make(map[int]map[int]bool)
		var relations [] db.Relation
		db.Mysql.Where(`(( source >= ? AND source < ? ) OR ( target >= ? AND target < ? ) ) AND removed = ?`,
			i, i+chunkSize, i, i+chunkSize, 0).Find(&relations)

		for _, rel := range relations {
			if sourceToTarget[rel.Source] == nil {
				sourceToTarget[rel.Source] = make(map[int]bool)
			}
			sourceToTarget[rel.Source][rel.Target] = true
		}

		for _, rel := range relations {
			if subMap, ok := sourceToTarget[rel.Target]; ok {
				if _, ok := subMap[rel.Source]; ok {
					continue
				}
			}
			relationIDNeedToRemove[rel.ID] = true
		}

		var ids = make([]string, 0, len(relationIDNeedToRemove))
		for key := range relationIDNeedToRemove {
			ids = append(ids, key)
		}
		if len(ids) != 0 {
			db.Mysql.Table(subjectTable).Where(`id IN (?)`, ids).Update(`locked`, 1)
		}
	}
}
func firstRun(subjectStart int, subjectEnd int) {
	var doneIDs = make(map[int]bool)

	var subjects = make(map[int]db.Subject)
	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var s []db.Subject
		db.Mysql.Where(`id >= ? AND id < ? AND locked = ? AND subject_type != ?`, i, i+chunkSize, 0, typeMusic).Find(&s)
		for _, subject := range s {
			if subject.SubjectType == typeMusic || subject.Locked != 0 {
				logrus.Errorf("subject error %v\n", subject.ID)
				continue
			}
			subject.Map = 0
			subjects[subject.ID] = subject
		}
	}
	logrus.Infof("total subject %d", len(subjects))

	var relationFromId = make(map[int]map[int]db.Relation)
	var edgeCount = 0

	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var edges []db.Relation
		db.Mysql.Table(relationTable).
			Where(`(source >= ?) AND (source < ?) AND (removed = ?)`, i, i+chunkSize, 0).Find(&edges)
		for _, edge := range edges {
			edgeCount += 1
			edge.Map = 0
			if relationFromId[edge.Source] == nil {
				relationFromId[edge.Source] = make(map[int]db.Relation)
			}
			relationFromId[edge.Source][edge.Target] = edge
			if relationFromId[edge.Target] == nil {
				relationFromId[edge.Target] = make(map[int]db.Relation)
			}
			relationFromId[edge.Target][edge.Source] = edge
		}
	}
	logrus.Infof("total %d edges", edgeCount)

	var dealWithNode func(int)
	dealWithNode = func(sourceID int) {
		var s, ok = subjects[sourceID]
		if ok {
			return
		}
		var edges = relationFromId[sourceID]
		var mapID = 0
		for _, edge := range edges {
			if edge.Map != 0 {
				mapID = edge.Map
				break
			}
		}

		if mapID == 0 {
			m.Lock()
			maxMapID++
			mapID = maxMapID
			m.Unlock()
		}
		for _, edge := range edges {
			edge.Map = mapID
		}
		s.Map = mapID
		doneIDs[sourceID] = true
		for _, edge := range edges {
			dealWithNode(edge.Target)
		}
		doneIDs[sourceID] = true
	}

	for _, subject := range subjects {
		dealWithNode(subject.ID)
	}

	logrus.Infoln(len(doneIDs))
	var subjectMaps = make(map[int][]int)
	var relationMaps = make(map[int][]string)
	for _, subject := range subjects {
		subjectMaps[subject.Map] = append(subjectMaps[subject.Map], subject.ID)
	}

	for _, edges := range relationFromId {
		for _, edge := range edges {
			relationMaps[edge.Map] = append(relationMaps[edge.Map], edge.ID)
		}
	}

	for key, ids := range subjectMaps {
		err := chunkIterInt(ids, func(s []int) error {
			return db.Mysql.Table(subjectTable).Where(`id IN (?)`, s).Update(`map`, key).Error
		})
		if err != nil {
			logrus.Errorln(err)
		}
	}

	for key, ids := range relationMaps {
		err := chunkIterStr(ids, func(s []string) error {
			return db.Mysql.Table(relationTable).Where(`id IN (?)`, s).Update(`map`, key).Error
		})
		if err != nil {
			logrus.Errorln(err)
		}
	}
	logrus.Infoln("finish save to db")
}

func chunkIterInt(s []int, f func([]int) error) error {

	var err error
	l := len(s)
	for i := 0; i < l; i += chunkSize {
		err = f(s[i:min(i+chunkSize, l)])
		if err != nil {
			return err
		}
	}
	return err
}

func chunkIterStr(s []string, f func([]string) error) error {

	var err error
	l := len(s)
	for i := 0; i < l; i += chunkSize {
		err = f(s[i:min(i+chunkSize, l)])
		if err != nil {
			return err
		}
	}
	return err
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

func Init() {
	reCalculateMap()
}
