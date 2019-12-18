package bgmTv

import (
	"bgm38/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jordic/goics"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func index(ctx *gin.Context) {
	ctx.String(200, "place holder")
}

var client = resty.New()
var cstZone = time.FixedZone("CST", 8*3600)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}
func userCalendar(ctx *gin.Context) {

	userId := ctx.Param("user_id")
	if userId == "" {
		ctx.String(401, "should give a userId")
		return
	}

	resp, err := client.R().
		SetQueryParam("cat", "watching").
		Get(fmt.Sprintf("https://mirror.api.bgm.rin.cat/user/%s/collection", userId))
	if err != nil {
		logrus.Debugln(err)
		ctx.String(502, "gateway error")
		return
	}

	var data []model.UserCollection

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Debugln(err)
		ctx.String(502, "gateway error")
		return
	}

	cal := goics.NewComponent()
	cal.SetType("VCALENDAR")
	cal.AddProperty("prodid", "-//trim21//www.trim21.cn//")
	cal.AddProperty("CALSCAL", "GREGORIAN")
	cal.AddProperty("PRODID;X-RICAL-TZSOURCE=TZINFO", "-//tmpo.io")

	cal.AddProperty("version", "2.0")
	cal.AddProperty("name", "Followed Bangumi Calendar")
	cal.AddProperty("description", "Followed Bangumi Calendar")
	cal.AddProperty("X-WR-CALNAM", "Followed Bangumi Calendar")
	cal.AddProperty("X-WR-CALDESC", "Followed Bangumi Calendar")

	for _, subject := range data {
		s := goics.NewComponent()

		s.SetType("VEVENT")
		offset_day := getAirDayOffset(time.Now().In(cstZone).Weekday(), subject.Subject.AirWeekday)
		dt := time.Now().In(cstZone).Add(time.Duration(offset_day*24) * time.Hour)
		k, v := goics.FormatDateField("DTEND", dt)
		s.AddProperty(k, v)
		k, v = goics.FormatDateField("DTSTART", dt)
		s.AddProperty(k, v)

		s.AddProperty("UID", fmt.Sprintf("%d", subject.SubjectId))
		s.AddProperty("SUMMARY", subject.Subject.NameCn)
		cal.AddComponent(s)
	}

	ctx.Status(http.StatusOK)
	ctx.Header("Content-type", "text/calendar")
	ctx.Header("charset", "utf-8")
	ctx.Header("Content-Disposition", "inline")
	ctx.Header("filename", "calendar.ics")
	cal.Write(goics.NewICalEncode(ctx.Writer))
}

func getAirDayOffset(w time.Weekday, airWeekday int) int {
	weekday := int(w)
	airWeekday = airWeekday % 7

	return (airWeekday + 7 - weekday) % 7
}
