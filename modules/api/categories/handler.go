package categories

import (
	"log"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/kataras/iris/v12"
)

type handler struct {
	service domains.CategoryService
}

func NewHandler(service domains.CategoryService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAll(ctx iris.Context) {
	name := ctx.URLParam("name")
	params := map[string]interface{}{
		"name": name,
	}

	resp, err := h.service.GetAll(ctx.Request().Context(), params)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}

func (h *handler) GetById(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		log.Println(err)
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp, err := h.service.GetById(ctx.Request().Context(), id)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}
