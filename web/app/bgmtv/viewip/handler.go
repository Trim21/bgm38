package viewip

import (
	"encoding/json"
	"net/http"

	"bgm38/web/app/res"
	"bgm38/web/app/db"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

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
func viewIP(ctx iris.Context, subjectID int) {
	var s db.Subject
	err := db.DB.Select("id, map").Where("id = ?", subjectID).Find(&s).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.StatusCode(404)
			_, err := ctx.JSON(res.Error{
				Message: "subject not found",
				Status:  "not_found",
			})
			if err != nil {
				logrus.Errorln(err)
			}
			return
		}
		logrus.Debugln(err)
		ctx.StatusCode(502)
		_, err := ctx.JSON(res.Error{
			Message: "can't tell you any about it",
			Status:  "error",
		})
		if err != nil {
			logrus.Errorln(err)
		}
		return
	}

	var relations []db.Relation

	mapID := s.Map
	if mapID == 0 {
		ctx.StatusCode(404)
		_, err = ctx.JSON(res.Error{
			Message: "subject doesn't have any related subjects",
			Status:  "not_found",
		})
		if err != nil {
			logrus.Errorln(err)
		}
		return
	}
	var subjects []db.Subject
	db.DB.Where("map = ?", mapID).Find(&relations)
	db.DB.Where("map = ?", mapID).Find(&subjects)
	ctx.StatusCode(200)
	_, err = ctx.JSON(subjectMapRes{
		Data:    *formatData(subjects, relations),
		Status:  http.StatusOK,
		Message: "nothing",
	})
	if err != nil {
		logrus.Errorln(err)
	}

}

type subject struct {
	db.Subject
	Tags         []db.Tag            `json:"tags"`
	Info         map[string][]string `json:"info"`
	ScoreDetails map[string]string   `json:"score_details"`
}

func getSubjectFullInfo(c iris.Context, subjectID int) {
	var s db.Subject
	err := db.DB.Where("id = ?", subjectID).Find(&s).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.StatusCode(404)
			c.Text("found nothing")
			return
		}
		logrus.Debugln(err)
		c.StatusCode(502)
		c.Text("233")
		return
	}

	var tags []db.Tag
	db.DB.Where("subject_id = ?", subjectID).Order("count desc").Find(&tags)

	var info map[string][]string
	_ = json.Unmarshal([]byte(s.Info), &info)

	var scoreDetails map[string]string
	_ = json.Unmarshal([]byte(s.ScoreDetails), &scoreDetails)
	c.StatusCode(200)
	c.JSON(subject{
		Subject:      s,
		Tags:         tags,
		Info:         info,
		ScoreDetails: scoreDetails,
	})
}
