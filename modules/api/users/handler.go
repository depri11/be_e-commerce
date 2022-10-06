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
	resp, err := h.service.GetAll(ctx.Request().Context())
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
	resp, err := h.service.GetByEmail(ctx.Request().Context(), email)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}

func (h *handler) Login(ctx iris.Context) {
	var payload input.UserLoginInput
	err := ctx.ReadJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp, err := h.service.Login(ctx.Request().Context(), &payload)
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

	resp, err := h.service.Register(ctx.Request().Context(), &payload)
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

	resp, err := h.service.Update(ctx.Request().Context(), email, &payload)
	if err != nil {
		resp.ResponseJSON(ctx)
		return
	}

	resp.ResponseJSON(ctx)
}
