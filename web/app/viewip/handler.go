package viewip

import (
	"encoding/json"

	"bgm38/web/app/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type info struct {
	SubjectID int `uri:"subject_id" binding:"required"`
}

func viewIP(c *gin.Context) {

	var v info
	if err := c.ShouldBindUri(&v); err != nil {
		c.JSON(400, gin.H{"msg": "subject_id should be int"})
		return
	}

	var s db.Subject
	err := db.DB.Select("id, map").Where("id = ?", v.SubjectID).Find(&s).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.String(404, "found nothing")
			return
		}
		logrus.Debugln(err)
		c.String(502, "233")
		return
	}

	var relations []db.Relation

	mapID := s.Map
	if mapID == 0 {
		c.JSON(404, gin.H{"message": "subject not found"})
		return
	}
	var subjects []db.Subject
	db.DB.Where("map = ?", mapID).Find(&relations)
	db.DB.Where("map = ?", mapID).Find(&subjects)
	c.JSON(200, formatData(subjects, relations))
}

type subject struct {
	db.Subject
	Tags         []db.Tag            `json:"tags"`
	Info         map[string][]string `json:"info"`
	ScoreDetails map[string]string   `json:"score_details"`
}

func getSubjectFullInfo(c *gin.Context) {

	var v info
	if err := c.ShouldBindUri(&v); err != nil {
		c.JSON(400, gin.H{"msg": "subject_id should be int"})
		return
	}

	var s db.Subject
	err := db.DB.Where("id = ?", v.SubjectID).Find(&s).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.String(404, "found nothing")
			return
		}
		logrus.Debugln(err)
		c.String(502, "233")
		return
	}

	var tags []db.Tag
	db.DB.Where("subject_id = ?", v.SubjectID).Order("count desc").Find(&tags)

	var info map[string][]string
	_ = json.Unmarshal([]byte(s.Info), &info)

	var scoreDetails map[string]string
	_ = json.Unmarshal([]byte(s.ScoreDetails), &scoreDetails)

	c.JSON(200, subject{
		Subject:      s,
		Tags:         tags,
		Info:         info,
		ScoreDetails: scoreDetails,
	})
}
