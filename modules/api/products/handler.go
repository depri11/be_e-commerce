package products

import (
	"log"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/kataras/iris/v12"
)

type handler struct {
	service domains.ProductService
}

func NewHandler(service domains.ProductService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAll(ctx iris.Context) {
	products, code, err := h.service.GetAll()
	if err != nil {
		ctx.StatusCode(code)
		log.Println(err)
		ctx.JSON(products)
		return
	}
	ctx.JSON(products)
}
