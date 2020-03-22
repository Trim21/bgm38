package docs

import "bgm38/config"

func OpenApi() string {
	SwaggerInfo.Version = config.Version
	return (&s{}).ReadDoc()
}
