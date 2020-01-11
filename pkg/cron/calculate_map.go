package cron

import (
	"fmt"
	"sync"

	"bgm38/pkg/db"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
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

func reCalculateMap() {
	var err error
	db.InitDB()
	var maxSubject db.Subject
	var minSubject db.Subject
	db.Mysql.Order(`id desc`).First(&maxSubject)
	db.Mysql.Order(`id`).First(&minSubject)

	tx := db.Mysql.Begin()

	if tx.Table(relationTable).Where("1 = ?", 1).
		Update(`removed`, 0).Error != nil {
	}
	preRemove(tx, minSubject.ID, maxSubject.ID)
	tx.Commit()

	tx = db.Mysql.Begin()
	err = firstRun(tx, minSubject.ID, maxSubject.ID)
	if err != nil {
		tx.Rollback()
		logrus.Errorln(err)
		return
	}
	tx.Commit()
}

func removeNodes(tx *gorm.DB, nodeIDs ...int) {
	tx.Table(subjectTable).Where(`id in (?)`, nodeIDs).
		Update(`locked`, 1)
	tx.Table(relationTable).
		Where(`target in (?) or source in (?)`, nodeIDs, nodeIDs).
		Update(`removed`, 1)
}

func relationsNeedToRemove(tx *gorm.DB, m map[int]int) {
	for id1, id2 := range m {
		tx.Table(relationTable).
			Where(`(target = ? AND source = ?) OR (target = ? AND source = ?)`,
				id1, id2, id2, id1).
			Update(`removed`, 1)

	}
}

func preRemove(tx *gorm.DB, subjectStart int, subjectEnd int) {
	println("pre remove")

	removeNodes(tx, 91493, 102098, 228714, 231982, 932, 84944, 78546)

	tx.Table(relationTable).
		Where(`relation in (?)`, blankList).Update(`removed`, 1)

	relationsNeedToRemove(tx, map[int]int{
		91493:  8,
		8108:   35866,
		446:    123207,
		123207: 27466,
		123217: 4294,
	})

	tx.Table(subjectTable).Where(`subject_type = ?`, typeMusic).
		Update(`locked`, 1)
	var idToRemove = make(map[int]bool)
	var subjects []db.Subject
	tx.Where(`locked = ?`, 1).Find(&subjects)
	for _, s := range subjects {
		idToRemove[s.ID] = true
	}

	var idToRM []int
	for key := range idToRemove {
		idToRM = append(idToRM, key)
	}

	err := chunkIterInt(idToRM, func(s []int) error {
		return tx.Table(relationTable).
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
		tx.Where(`(( source >= ? AND source < ? ) `+
			`OR ( target >= ? AND target < ? ) ) AND removed = ?`,
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
			tx.Table(relationTable).Where(`id IN (?)`, ids).
				Update(`removed`, 1)
		}
	}
	println("finish pre remove")
}

func getRelationsFromDB(tx *gorm.DB, subjectStart, subjectEnd int) (map[int]map[string]*db.Relation, int, error) {

	edgeCount := 0
	var relationFromId = make(map[int]map[string]*db.Relation)
	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var edges = make([]db.Relation, 0, 5000)

		err := tx.Table(relationTable).
			Where(`(source >= ?) AND (source < ?) AND (removed = ?)`, i, i+chunkSize, 0).
			Find(&edges).Error
		if err != nil {
			return nil, 0, err
		}

		for _, edge := range edges {
			if edge.Removed == 1 {
				continue
			}
			edge.Map = 0
			edgeCount += 1
			if relationFromId[edge.Source] == nil {
				relationFromId[edge.Source] = make(map[string]*db.Relation)
			}

			if relationFromId[edge.Target] == nil {
				relationFromId[edge.Target] = make(map[string]*db.Relation)
			}

			var edgeCopy = db.Relation{}

			err := copier.Copy(&edgeCopy, &edge)
			if err != nil {
				return nil, 0, err
			}

			relationFromId[edge.Source][edge.ID] = &edgeCopy
			relationFromId[edge.Target][edge.ID] = &edgeCopy
		}
	}

	return relationFromId, edgeCount, nil
}
func getSubjectsFromDB(tx *gorm.DB, subjectStart int, subjectEnd int) (map[int]*db.Subject, error) {
	var subjects = make(map[int]*db.Subject)
	for i := subjectStart; i < subjectEnd; i += chunkSize {
		var s = make([]db.Subject, 0, 5000)

		err := tx.Where(`id >= ? AND id < ? AND `+
			`locked = ? AND subject_type != ?`, i, i+chunkSize, 0, typeMusic).
			Find(&s).Error
		if err != nil {
			return nil, err
		}

		for _, subject := range s {
			if subject.SubjectType == typeMusic || subject.Locked != 0 {
				logrus.Errorf("subject error %v\n", subject.ID)
				continue
			}
			subject.Map = 0
			subjectCopy := db.Subject{}
			err := copier.Copy(&subjectCopy, &subject)
			if err != nil {
				return nil, err
			}

			subjects[subject.ID] = &subjectCopy
		}
		s = nil
	}
	return subjects, nil
}

func firstRun(tx *gorm.DB, subjectStart int, subjectEnd int) error {
	logrus.Debugf("build relation map with start id %d and end id %d", subjectStart, subjectEnd)
	var doneIDs = make(map[int]bool, subjectEnd-subjectStart)
	subjects, err := getSubjectsFromDB(tx, subjectStart, subjectEnd)
	if err != nil {
		return err
	}

	logrus.Infof("total subject %d", len(subjects))

	relationFromId, edgeCount, err := getRelationsFromDB(tx, subjectStart, subjectEnd)

	if err != nil {
		return err
	}

	if m, ok := relationFromId[8108]; !ok {
		logrus.Fatalln("err")
	} else {
		for id, edge := range m {
			if edge.Removed == 1 {
				fmt.Println(id, edge)
			}
		}
	}

	logrus.Infof("total %d edges", edgeCount)
	logrus.Debugf("id 8 get %d edges", len(relationFromId[8]))

	count := 0
	var dealWithNode func(int)
	dealWithNode = func(sourceID int) {
		if _, ok := doneIDs[sourceID]; ok {
			return
		}
		count++
		var edges = relationFromId[sourceID]
		var mapID = 0
		for _, edge := range edges {
			if edge.Map != 0 {
				mapID = edge.Map
				break
			}
			if subjects[edge.Target].Map != 0 {
				mapID = subjects[edge.Target].Map
				break
			}
			if subjects[edge.Source].Map != 0 {
				mapID = subjects[edge.Source].Map
				break
			}
		}

		if mapID == 0 {
			m.Lock()
			maxMapID++
			mapID = maxMapID
			m.Unlock()
		}
		for k := range edges {
			edges[k].Map = mapID
		}
		subjects[sourceID].Map = mapID
		doneIDs[sourceID] = true
		for _, edge := range edges {
			dealWithNode(edge.Target)
		}
	}

	logrus.Debugln("now iter %d subjects", len(subjects))

	for id := range subjects {
		dealWithNode(id)
	}
	for key, subject := range subjects {
		if key != subject.ID {
			logrus.Fatalf("%s %s", key, subject.ID)
		}
	}
	logrus.Debugf("called %d times", count)
	logrus.Debugf("done %d ids", len(doneIDs))

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

	logrus.Infof("got %d map", len(subjectMaps))
	for key, ids := range subjectMaps {
		logrus.Debugln(key)
		err := chunkIterInt(ids, func(s []int) error {
			return tx.Table(subjectTable).Where(`id IN (?)`, s).
				Update(`map`, key).Error
		})
		if err != nil {
			return nil
		}
	}

	for key, ids := range relationMaps {
		logrus.Debugln(key)
		err := chunkIterStr(ids, func(s []string) error {
			return tx.Table(relationTable).Where(`id IN (?)`, s).
				Update(`map`, key).Error
		})
		if err != nil {
			return err
		}
	}

	tx.Table(subjectTable).Where(`locked = ?`, 1).
		Update(`map`, 0)

	tx.Table(relationTable).Where(`removed = ?`, 1).
		Update(`map`, 0)

	logrus.Infoln("finish save to db")
	return nil
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

func Init() {
	reCalculateMap()
}
