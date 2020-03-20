package spider

import (
	"bgm38/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var relationUpsertStmt *sqlx.NamedStmt
var tagUpsertStmt *sqlx.NamedStmt
var subjectUpsertStmt *sqlx.NamedStmt
var epUpsertStmt *sqlx.NamedStmt

func prepareStmt() {

	var err error

	if relationUpsertStmt == nil {
		relationUpsertStmt, err = db.MysqlX.PrepareNamed("INSERT INTO `relation` (`id`, `relation`, `source`, `target`) VALUES (:id, :relation, :source, :target) ON DUPLICATE KEY UPDATE `relation` = VALUES(`relation`)")
		if err != nil {
			logrus.Fatalln("prepare statements error for relation upsert", err)
		}
	}

	if tagUpsertStmt == nil {
		tagUpsertStmt, err = db.MysqlX.PrepareNamed("INSERT INTO `tag` (`subject_id`, `text`, `count`) VALUES (:subject_id, :text, :count) ON DUPLICATE KEY UPDATE `count` = VALUES(`count`)")
		if err != nil {
			logrus.Fatalln("prepare statements error for tag upsert", err)
		}
	}
	if subjectUpsertStmt == nil {
		raw := `INSERT INTO subject (id, name, image, subject_type, name_cn, tags, info, score_details, score, wishes, done, doings, on_hold, dropped, map, locked) VALUES (:id,:name,:image,:subject_type,:name_cn,:tags,:info,:score_details,:score,:wishes,:done,:doings,:on_hold,:dropped,:map,:locked) ON DUPLICATE KEY UPDATE  name = :name, image = :image, subject_type = :subject_type, name_cn = :name_cn, tags = :tags, info = :info, score_details = :score_details, score = :score, wishes = :wishes, done = :done, doings = :doings, on_hold = :on_hold, dropped = :dropped, locked = :locked`
		subjectUpsertStmt, err = db.MysqlX.PrepareNamed(raw)
		if err != nil {
			logrus.Fatalln("prepare statements error for tag upsert", err)
		}
	}
	if epUpsertStmt == nil {
		epUpsertStmt, err = db.MysqlX.PrepareNamed("INSERT INTO `ep`(`ep_id`, `subject_id`, `name`, `episode`) VALUES (:ep_id, :subject_id, :name, :episode) ON DUPLICATE KEY UPDATE `name` = VALUES(`name`), `episode` = VALUES(`episode`)")
		if err != nil {
			logrus.Fatalln("prepare statements error for ep upsert", err)
		}
	}
}
