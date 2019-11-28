package setup

import (
	"github.com/kataras/iris"
	"bgm38/pkg/server/bindata"
	"os"
)

func Template(app *iris.Application) {
	engine := iris.HTML("./bindata/templates", ".tmpl")
	if os.Getenv("DEV") == "" {
		// production env
		engine = engine.Binary(bindata.Asset, bindata.AssetNames)
	} else {
		engine.Reload(true)
	}
	app.RegisterView(engine)
}

func Static(app *iris.Application) {
	if os.Getenv("DEV") == "" {
		app.StaticEmbeddedGzip("/static/",
			"./bindata/gzipped/static/",
			bindata.GzipAsset, bindata.GzipAssetNames)
	} else {
		app.StaticWeb("/static/", "./bindata/gzipped/static/")
	}
}
