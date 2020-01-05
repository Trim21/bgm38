package cron

import (
	"fmt"

	"bgm38/pkg/db"
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

	// id_to_remove = []
	// Subject.update(locked = 1).where(Subject.subject_type == 'Music').execute()
	// for s
	// 	in
	// Subject.
	// select (Subject.id).where(Subject.locked == 1):
	// 	id_to_remove.append(s.id)
	// 	Relation.update(
	// 		removed = 1
	// 	).where(Relation.source.in_(id_to_remove) |
	// 		Relation.target.in_(id_to_remove)).execute()
	//
	// 	for chunk
	// 		in
	// 	chunk_iter_list(list(range(subject_start, subject_end))):
	// 	db_data = list(
	// 		Subject.
	// 	select (
	// 		Subject.id,
	// 		Subject.subject_type,
	// 		Subject.locked,
	// 	).where(
	// 			Subject.id.in_(chunk) & (Subject.subject_type != 'Music') &
	// 				(Subject.locked == 0)
	// 		)
	// 	)
	// 		for s
	// 			in
	// 		db_data:
	// 		assert
	// 		s.subject_type != 'Music'
	// 		assert
	// 		s.locked == 0
	// 		non_exists_ids = list(set(chunk) -
	// 		{
	// 			x.id
	// 			for x
	// 				in
	// 			db_data
	// 		})
	// 		Relation.update(removed = 1).where(
	// 			Relation.source.in_(non_exists_ids) | Relation.target.in_(non_exists_ids)
	// 		).execute()
	//
	// 		for i
	// 		in range
	// 		(subject_start, subject_end, CHUNK_SIZE):
	// 		relation_id_need_to_remove = set()
	// 	source_to_target:
	// 		Dict
	// 		[int, Dict] = defaultdict(dict)
	// 		sources = Relation.select ().where((((Relation.source >= i) &
	// 		(Relation.source < i + CHUNK_SIZE)) |
	// 		((Relation.target >= i) &
	// 		(Relation.target < i + CHUNK_SIZE))) &
	// 		(Relation.removed == 0))
	//
	// 		sources = list(sources)
	//
	// 		for edge in sources:
	// 		source_to_target[edge.source][edge.target] = True
	//
	// 		for edge in sources:
	// 		if not source_to_target[edge.target].get(edge.source):
	// 		relation_id_need_to_remove.add(edge.id)
	//
	// 		for chunk in chunk_iter_list(list(relation_id_need_to_remove)):
	// 		Relation.update(removed = 1).where(Relation.id.in_(chunk)).execute()
	// 		print('finish pre remove')
	//
	//
}
