package viewip

import (
	"encoding/json"
	"net/http"

	"bgm38/web/app/db"
	"bgm38/web/app/res"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// swagger:parameters viewIP
type info struct {
	// bgm.tv subject id
	// in: path
	SubjectID int `uri:"subject_id" binding:"required" json:"subject_id"`
}

// @ID viewIP
// @Summary get subject relations map
// @Description Get related Subjects and their relations
// @Produce  json
// @Param subject_id path int true "Account ID"
// @Success 200 {object} subjectMapRes
// @Failure 422 {object} res.ValidationError
// @Failure 404 {object} res.Error
// @Failure 500 {object} res.Error
// @Router /bgm.tv/v2/view_ip/subject/{subject_id} [get]
func viewIP(c *gin.Context) {

	var v info
	if err := c.ShouldBindUri(&v); err != nil {
		c.JSON(422, res.ValidationError{
			Message:   "subject_id should be int",
			FieldName: "subject_id",
		})
		return
	}

	var s db.Subject
	err := db.DB.Select("id, map").Where("id = ?", v.SubjectID).Find(&s).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(404, res.Error{
				Message: "subject not found",
				Status:  "not_found",
			})
			return
		}
		logrus.Debugln(err)
		c.JSON(502, res.Error{
			Message: "can't tell you any about it",
			Status:  "error",
		})
		return
	}

	var relations []db.Relation

	mapID := s.Map
	if mapID == 0 {
		c.JSON(404, res.Error{
			Message: "subject doesn't have any related subjects",
			Status:  "not_found",
		})
		return
	}
	var subjects []db.Subject
	db.DB.Where("map = ?", mapID).Find(&relations)
	db.DB.Where("map = ?", mapID).Find(&subjects)
	c.JSON(200, subjectMapRes{
		Data:    *formatData(subjects, relations),
		Status:  http.StatusOK,
		Message: "nothing",
	})
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
