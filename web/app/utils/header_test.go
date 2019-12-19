package utils_test

import (
	"bgm38/web/app/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsBrowser(t *testing.T) {
	buf := new(bytes.Buffer)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", buf)
	for _, header := range []string{
		"Mozilla/5.0 (Windows NT 6.1; rv:12.0) Gecko/20120403211507 Firefox/12.0",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (Linux; Android 7.0; MI 5s Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.145 Mobile Safari/537.36 EdgA/41.0.0.1273",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/601.5.17 (KHTML, like Gecko)",
	} {
		c.Request.Header.Set("user-agent", header)
		assert.True(t, utils.IsBrowser(c), fmt.Sprintf("user-agent %s should be consider as browser", header))

	}

	for _, header := range []string{
		"Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36; 360Spider	",
		"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"iOS/11.4.1 (15G77) dataaccessd/1.0", // ios calendar
		"http.rb/4.2.0",
		"Go-http-client/1.1",
		"okhttp/3.10.0",
	} {
		c.Request.Header.Set("user-agent", header)
		assert.False(t, utils.IsBrowser(c), fmt.Sprintf("user-agent %s should not be consider as browser", header))

	}

}
