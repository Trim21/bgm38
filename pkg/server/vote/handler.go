package vote

import (
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	//ctx.String(200, "hello world")
	ctx.HTML(200, "vote/index.tmpl", map[string]string{
		"Title":   "hello Vote Function",
		"Message": "hello world",
	})
}

