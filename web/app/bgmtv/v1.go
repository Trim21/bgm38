package bgmtv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bgm38/pkg/model"
	"bgm38/web/app/res"
	"bgm38/web/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jordic/goics"
	"github.com/sirupsen/logrus"
)

func index(ctx *gin.Context) {
	ctx.String(200, "place holder")
}

var client = resty.New()
var cstZone = time.FixedZone("CST", 8*3600)

// @ID watchingCalendarV1
// @Summary generate a calendar from watching collection
// @Description 根据在看的番剧生成ics格式日历
// @Description 如果浏览器访问时会返回纯文本数据
// @Description 在使用日历app导入时会返回日历数据
// @Produce  text/calendar
// @Produce  json
// @Param user_id path string true "user_id, 可以在个人主页的网址中找到"
// @Success 200 {string} string "text/calendar"
// @Failure 422 {object} res.ValidationError
// @Failure 404 {object} res.Error
// @Failure 502 {object} res.Error
// @Router /bgm.tv/v1/calendar/{user_id} [get]
func userCalendar(ctx *gin.Context) {

	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(422, res.ValidationError{
			Message:   "need a user_id",
			FieldName: "user_id",
		})
		return
	}
	resp, err := client.R().SetQueryParam("cat", "watching").
		Get(fmt.Sprintf("https://mirror.api.bgm.rin.cat/user/%s/collection", userID))
	if err != nil {
		logrus.Debugln(err)
		ctx.JSON(502, res.Error{
			Message: "connect to mirror.api.bgm.rin.cat error",
			Status:  "error",
		})
		return
	}

	var data []model.UserCollection

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Debugln(err)
		ctx.JSON(404, res.Error{
			Message: "User doesn't exist",
			Status:  "error",
		})
		return
	}

	cal := goics.NewComponent()
	cal.SetType("VCALENDAR")
	cal.AddProperty("prodid", "-//trim21//www.trim21.cn//")
	cal.AddProperty("CALSCAL", "GREGORIAN")
	cal.AddProperty("PRODID;X-RICAL-TZSOURCE=TZINFO", "-//tmpo.io")
	cal.AddProperty("version", "2.0")
	cal.AddProperty("name", "Bgm.tv Followed Bangumi Calendar")
	cal.AddProperty("description", "Followed Bangumi Calendar")
	cal.AddProperty("X-WR-CALNAM", "Followed Bangumi Calendar")
	cal.AddProperty("X-WR-CALDESC", "Followed Bangumi Calendar")

	for _, subject := range data {
		s := goics.NewComponent()

		s.SetType("VEVENT")
		offsetDay := getAirDayOffset(time.Now().In(cstZone).Weekday(), subject.Subject.AirWeekday)
		dt := time.Now().In(cstZone).Add(time.Duration(offsetDay*24) * time.Hour)
		k, v := goics.FormatDateField("DTEND", dt)
		s.AddProperty(k, v)
		k, v = goics.FormatDateField("DTSTART", dt)
		s.AddProperty(k, v)

		s.AddProperty("UID", fmt.Sprintf("%d", subject.SubjectID))
		s.AddProperty("SUMMARY", subject.Subject.NameCn)
		cal.AddComponent(s)
	}

	ctx.Status(http.StatusOK)

	if !utils.IsBrowser(ctx) {
		ctx.Header("charset", "utf-8")
		ctx.Header("Content-type", "text/calendar")
		ctx.Header("Content-Disposition", "inline")
		ctx.Header("filename", "calendar.ics")
	}

	cal.Write(goics.NewICalEncode(ctx.Writer))
}

func getAirDayOffset(w time.Weekday, airWeekday int) int {
	weekday := int(w)
	airWeekday = airWeekday % 7

	return (airWeekday + 7 - weekday) % 7
}
