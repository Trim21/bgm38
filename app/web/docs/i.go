package docs

import "bgm38/config"

// OpenAPI
func OpenAPI() string {
	SwaggerInfo.Version = config.Version
	return (&s{}).ReadDoc()
}
