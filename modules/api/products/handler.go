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
	order := ctx.URLParam("order_by")
	sort := ctx.URLParam("sort_by")
	params := map[string]interface{}{
		"order_by": order,
		"sort_by":  sort,
	}

	res, err := h.service.GetAll(params)
	if err != nil {
		log.Println(err)
		res.ResponseJSON(ctx)
		return
	}
	res.ResponseJSON(ctx)
}

func (h *handler) GetByID(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		log.Println(err)
		ctx.StopWithError(500, err)
		return
	}

	res, err := h.service.GetById(id)
	if err != nil {
		log.Println(err)
		res.ResponseJSON(ctx)
		return
	}

	res.ResponseJSON(ctx)
}
