package router

import (
	"github.com/depri11/be_e-commerce/common/configs"
	"github.com/depri11/be_e-commerce/middleware"
	"github.com/depri11/be_e-commerce/modules/api/categories"
	"github.com/depri11/be_e-commerce/modules/api/products"
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

	userRepo := users.NewRepository(config.PostgreConfig.GormConn)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)
	noAuth.Get("/users", userHandler.GetAll)

	productRepo := products.NewRepository(config.PostgreConfig.GormConn)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)
	noAuth.Get("/products", productHandler.GetAll)
	noAuth.Get("/product", productHandler.GetByID)

	categoryRepo := categories.NewRepository(config.PostgreConfig.GormConn)
	categoryService := categories.NewService(categoryRepo)
	categoryHandler := categories.NewHandler(categoryService)
	noAuth.Get("/categories", categoryHandler.GetAll)
	noAuth.Get("/categorie", categoryHandler.GetById)

	return app
}
