package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func IsBrowser(ctx *gin.Context) bool {
	userAgent := strings.ToLower(ctx.GetHeader("user-agent"))
	if strings.Contains(userAgent, "mozilla") ||
		strings.Contains(userAgent, "chrome") ||
		strings.Contains(userAgent, "safari") {
		return true

	}
	return false
}
