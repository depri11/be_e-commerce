package deliveries

import (
	"github.com/depri11/be_e-commerce/domains"
	"github.com/kataras/iris/v12"
)

type deliveries struct {
	usecases domains.UserUsecases
}

func NewDeliveries(usecases domains.UserUsecases) *deliveries {
	return &deliveries{
		usecases: usecases,
	}
}

func (d *deliveries) GetAll(ctx iris.Context) {
	res, err := d.usecases.GetAll()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(res)
}
