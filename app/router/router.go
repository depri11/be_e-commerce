package router

import (
	"github.com/depri11/be_e-commerce/common/configs"
	"github.com/depri11/be_e-commerce/middleware"
	"github.com/depri11/be_e-commerce/modules/api/users"
	"github.com/kataras/iris/v12"
)

func Setup(config *configs.Configuration) *iris.Application {
	app := iris.New()

	m := middleware.NewMiddleware(config.RedisClient, config.JwtKey)
	app.Use(m.CORS())

	noAuth := app.Party("api/v1")

	auth := app.Party("api/v1")
	auth.Use(m.Customize())

	repo := users.NewRepository(config.PostgreConfig.GormConn)
	usecases := users.NewService(repo)
	deliveries := users.NewHandler(usecases)

	noAuth.Get("/", deliveries.GetAll)

	return app
}
