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
	userService := users.NewService(userRepo, config.TimeoutCtx)
	userHandler := users.NewHandler(userService)
	noAuth.Post("/register", userHandler.Register)
	noAuth.Post("/login", userHandler.Login)
	noAuth.Get("/users", userHandler.GetAll)
	noAuth.Get("/me", userHandler.MyProfile)
	noAuth.Put("/me/update", userHandler.EditProfile)

	productRepo := products.NewRepository(config.PostgreConfig.GormConn)
	productService := products.NewService(productRepo, config.TimeoutCtx)
	productHandler := products.NewHandler(productService)
	noAuth.Get("/products", productHandler.GetAll)
	noAuth.Get("/product", productHandler.GetByID)

	categoryRepo := categories.NewRepository(config.PostgreConfig.GormConn)
	categoryService := categories.NewService(categoryRepo, config.TimeoutCtx)
	categoryHandler := categories.NewHandler(categoryService)
	noAuth.Get("/categories", categoryHandler.GetAll)
	noAuth.Get("/categorie", categoryHandler.GetById)

	return app
}
