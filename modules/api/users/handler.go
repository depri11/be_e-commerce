package users

import (
	"log"

	"github.com/depri11/be_e-commerce/domains"
	"github.com/depri11/be_e-commerce/input"
	"github.com/kataras/iris/v12"
)

type handler struct {
	service domains.UserService
}

func NewHandler(service domains.UserService) *handler {
	return &handler{service}
}

func (h *handler) GetAll(ctx iris.Context) {
	resp, err := h.service.GetAll()
	if err != nil {
		log.Println(err)
		resp.ResponseJSON(ctx)
		return
	}
	resp.ResponseJSON(ctx)
}

func (h *handler) MyProfile(ctx iris.Context) {
	email := ctx.Values().GetString("user_email")
	log.Println(email)
	resp, err := h.service.GetByEmail(email)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}

func (h *handler) Register(ctx iris.Context) {
	var payload input.UserRegisterInput
	err := ctx.ReadJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp, err := h.service.Register(&payload)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}

func (h *handler) EditProfile(ctx iris.Context) {
	ctx.Values().Set("user_email", "dev@gmail.com")

	email := ctx.Values().GetString("user_email")
	log.Println(email)
	var payload input.UserEditProfileInput
	err := ctx.ReadJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp, err := h.service.Update(email, &payload)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}
