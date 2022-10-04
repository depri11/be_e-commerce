package users

import (
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
	res, err := d.usecases.GetAll()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(res)
}
