package router

import (
	"github.com/depri11/be_e-commerce/common/configs"
	"github.com/depri11/be_e-commerce/deliveries"
	"github.com/depri11/be_e-commerce/middleware"
	"github.com/depri11/be_e-commerce/repositories"
	"github.com/depri11/be_e-commerce/usecases"
	"github.com/kataras/iris/v12"
)

func Setup(config *configs.Configuration) *iris.Application {
	app := iris.New()

	m := middleware.NewMiddleware(config.RedisClient, config.JwtKey)
	app.Use(m.CORS())

	noAuth := app.Party("api/v1")

	auth := app.Party("api/v1")
	auth.Use(m.Customize())

	repo := repositories.NewRepository(config.PostgreConfig.GormConn)
	usecases := usecases.NewUsecases(repo)
	deliveries := deliveries.NewDeliveries(usecases)

	noAuth.Get("/", deliveries.GetAll)

	return app
}
