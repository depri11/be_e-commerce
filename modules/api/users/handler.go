package users

import (
	"log"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/kataras/iris/v12"
)

type handler struct {
	usecases domains.UserUsecases
}

func NewHandler(usecases domains.UserUsecases) *handler {
	return &handler{
		usecases: usecases,
	}
}

func (d *handler) GetAll(ctx iris.Context) {
	res, code, err := d.usecases.GetAll()
	if err != nil {
		ctx.StatusCode(code)
		log.Println(err)
		ctx.JSON(res)
		return
	}
	ctx.JSON(res)
}
