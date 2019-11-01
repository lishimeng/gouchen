package webserver

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/lishimeng/go-libs/web"
	"github.com/lishimeng/gouchen/internal/etc"
	"github.com/lishimeng/gouchen/internal/static"
)

func Run(components ...web.Component) {

	bs, err := static.Asset("index.html")
	indexHtml := ""
	if err == nil {
		indexHtml = string(bs)
	}
	web.New(web.ServerConfig{Listen: etc.Config.Web.Listen}).
		SetHomePage(indexHtml).
		AdvancedConfig(func(app *iris.Application) {
			app.Logger().SetLevel("info")
			app.Use(logger.New())
			app.OnErrorCode(404, func(ctx iris.Context) {
				_, _ = ctx.Writef("404[not found]")
			})
			app.StaticEmbedded("/", "", static.Asset, static.AssetNames)
		}).
		Start(components...)


}
