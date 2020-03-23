package header

import (
	"strings"
)

func IsUABrowser(userAgent string) bool {
	userAgent = strings.ToLower(userAgent)
	if strings.Contains(userAgent, "spider") ||
		strings.Contains(userAgent, "bot") {
		return false
	}
	if strings.Contains(userAgent, "mozilla") ||
		strings.Contains(userAgent, "chrome") ||
		strings.Contains(userAgent, "safari") {
		return true

	}
	return false
}
