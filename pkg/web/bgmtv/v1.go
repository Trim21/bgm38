package bgmtv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber"
	"github.com/jordic/goics"
	"go.uber.org/zap"

	"bgm38/config"
	"bgm38/pkg/model"
	"bgm38/pkg/utils"
	"bgm38/pkg/web/res"
	"bgm38/pkg/web/utils/header"
)

var client = resty.New()
var cstZone = time.FixedZone("CST", 8*3600)

// @ID watchingCalendarV1
// @Summary 在看番剧日历
// @Description 根据在看的番剧生成ics格式日历
// @Description 如果浏览器访问时会返回纯文本数据
// @Description 在使用日历app导入时会返回日历数据
// @Produce  text/calendar
// @Param user_id path string true "user_id, 可以在个人主页的网址中找到"
// @Success 200 {string} string "text/calendar"
// @Failure 422 {object} res.Error
// @Failure 404 {object} res.Error
// @Failure 502 {object} res.Error
// @Router /bgm.tv/v1/calendar/{user_id} [get]
func userCalendar(ctx *fiber.Ctx, logger *zap.Logger) error {
	userID := ctx.Params("user_id")
	if userID == "" {
		ctx.Status(422)
		return ctx.JSON(res.Error{
			Message: "missing `user_id` in path",
			Status:  "error",
		})
	}
	resp, err := client.R().SetQueryParam("cat", "watching").
		Get(fmt.Sprintf("https://mirror.api.bgm.rin.cat/user/%s/collection", userID))
	var netErr net.Error
	if err != nil {
		ctx.Status(502)
		if errors.As(err, &netErr) {
			if netErr.Timeout() {
				return ctx.JSON(res.Error{
					Message: "request to mirror.api.bgm.rin.cat timeout",
					Status:  "error",
				})
			}
		}
		return ctx.JSON(res.Error{
			Message: "request to mirror.api.bgm.rin.cat timeout",
			Status:  "error",
		})

	}

	var data []model.UserCollection

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logger.Debug(err.Error())
		ctx.Status(404)
		return ctx.JSON(res.Error{
			Message: "User doesn't exist or can't fetch data from upstream serer, try refresh your page",
			Status:  "error",
		})
	}
	cal := makeCal(userID, data)
	ctx.Status(http.StatusOK)

	if !header.IsUABrowser(ctx.Get("user-agent")) {
		ctx.Set("charset", "utf-8")
		ctx.Set("Content-type", "text/calendar")
		ctx.Set("Content-Disposition", "inline")
		ctx.Set("filename", "calendar.ics")
	}
	cal.Write(goics.NewICalEncode(ctx.Fasthttp.Response.BodyWriter()))
	return nil
}

func getAirDayOffset(w time.Weekday, airWeekday int) int {
	weekday := int(w)
	airWeekday %= 7

	return (airWeekday + 7 - weekday) % 7
}

func makeCal(userID string, data []model.UserCollection) *goics.Component {
	cal := goics.NewComponent()
	cal.SetType("VCALENDAR")
	cal.AddProperty("name", "Bgm.tv Followed Bangumi Calendar")
	cal.AddProperty("description", utils.StrConcat(userID, " Followed Bangumi Calendar"))
	cal.AddProperty("PRODID", utils.StrConcat("-//trim21//api.bgm38.com//", config.Version, "//"))
	cal.AddProperty("X-WR-CALNAME", "bgm.tv")
	cal.AddProperty("version", "2.0")

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
	return cal
}
