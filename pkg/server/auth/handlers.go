package auth

import (
	"bgm38/config"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/url"
)

func callback(ctx *gin.Context) {
	ctx.String(200, "place holder")
}

var callbackURL = getCallbackURL()
var oauthURL = getOauthURL(callbackURL)

func redirect(ctx *gin.Context) {
	ctx.Redirect(307, oauthURL)
}

func getCallbackURL() string {
	return fmt.Sprintf("%s://%s/auth/v1/callback", config.PROTOCOL, config.VirtualHost)
}

func getOauthURL(callback string) string {
	u, err := url.Parse(`https://bgm.tv/oauth/authorize`)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("q", "golang")
	q.Set("client_id", config.AppID)
	q.Set("response_type", "code")
	q.Set("redirect_uri", callback)
	u.RawQuery = q.Encode()
	return u.String()
}
