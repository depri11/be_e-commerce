package router

import (
	"github.com/depri11/be_e-commerce/common/configs"
	"github.com/kataras/iris/v12"
)

func Setup(config *configs.Configuration) *iris.Application {
	app := iris.New()

	api := app.Party("api/v1")
	api.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, world!")
	})

	return app
}
