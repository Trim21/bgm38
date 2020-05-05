package vote

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"bgm38/pkg/fetch"
)

func TestRouterErrInput(t *testing.T) {
	app := fiber.New()
	Group(app.Group("/t"))
	req, _ := http.NewRequest("GET", "http://example.com/t/vote/a", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, 401, res.StatusCode)
	assert.Contains(t, string(body), "`topic_id` should be int", "hello")
}

func TestRouter(t *testing.T) {
	httpmock.ActivateNonDefault(fetch.GetClient().GetClient())

	defer httpmock.DeactivateAndReset()
	f, err := os.Open("../../../../tests/fixtures/vote/input.html")
	assert.Nil(t, err, "read fixtures error")
	body, _ := ioutil.ReadAll(f)
	httpmock.RegisterResponder("GET", "https://mirror.bgm.rin.cat/group/topic/1",
		httpmock.NewBytesResponder(200, body))

	app := fiber.New()
	Group(app.Group("/t"))
	req, _ := http.NewRequest("GET", "http://example.com/t/vote/1", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, 200, "should resp 200")
	body, _ = ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "hello", "hello")
}
