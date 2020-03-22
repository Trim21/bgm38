package bgmtv

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://127.0.0.1:3002"

func TestUserWatchingCalendar(t *testing.T) {
	httpmock.ActivateNonDefault(client.GetClient())

	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mirror.api.bgm.rin.cat/user/trim21/collection?cat=watching",
		httpmock.NewStringResponder(200, `[{
    "ep_status": 8,
    "lasttouch": 1584624657,
    "name": "Star Trek: Picard Season 1",
    "subject": {
      "air_date": "2020-01-23",
      "air_weekday": 4,
      "collection": {
        "doing": 8
      },
      "eps": 10,
      "eps_count": 10,
      "id": 298728,
      "images": {
        "common": "http://lain.bgm.tv/pic/cover/c/8c/0e/298728_y5UDL.jpg",
        "grid": "http://lain.bgm.tv/pic/cover/g/8c/0e/298728_y5UDL.jpg",
        "large": "http://lain.bgm.tv/pic/cover/l/8c/0e/298728_y5UDL.jpg",
        "medium": "http://lain.bgm.tv/pic/cover/m/8c/0e/298728_y5UDL.jpg",
        "small": "http://lain.bgm.tv/pic/cover/s/8c/0e/298728_y5UDL.jpg"
      },
      "name": "Star Trek: Picard Season 1",
      "name_cn": "星际迷航：皮卡德 第一季",
      "summary": "",
      "type": 6,
      "url": "http://bgm.tv/subject/298728"
    },
    "subject_id": 298728,
    "vol_status": 0
  }]`))

	app := fiber.New()
	Group(app)
	req, _ := http.NewRequest("GET", baseURL+"/bgm_tv/v1/calendar/trim21", nil)
	res, _ := app.Test(req)
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, 200, "should resp 200")
	assert.Contains(t, res.Header.Get("content-type"), "calendar",
		"response head should be calendar")
	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "皮卡德", "res should contains series name")
}
