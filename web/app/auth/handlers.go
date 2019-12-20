package auth

import (
	"fmt"
	"net/url"

	"bgm38/config"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func callback(ctx iris.Context) {
	ctx.StatusCode(200)
	ctx.Text("place holder")
}

var callbackURL = getCallbackURL()
var oauthURL = getOauthURL(callbackURL)

func redirect(ctx iris.Context) {
	ctx.Redirect(oauthURL, 307)
}

func getCallbackURL() string {
	return fmt.Sprintf("%s://%s/auth/v1/bgmtv.tv/callback", config.PROTOCOL, config.VirtualHost)
}

func getOauthURL(callback string) string {
	u, err := url.Parse(`https://bgmtv.tv/oauth/authorize`)
	if err != nil {
		logrus.Fatal(err)
	}
	q := u.Query()
	q.Set("q", "golang")
	q.Set("client_id", config.AppID)
	q.Set("response_type", "code")
	q.Set("redirect_uri", callback)
	u.RawQuery = q.Encode()
	return u.String()
}
