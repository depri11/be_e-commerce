package helpers

import "github.com/kataras/iris/v12"

type Response struct {
	Status  int         `json:"status" default:"200"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Other   interface{} `json:"other,omitempty"`
}

func (r *Response) ResponseJSON(ctx iris.Context) error {
	ctx.StatusCode(r.Status)
	return ctx.JSON(Response{Status: r.Status, Message: r.Message, Data: r.Data, Other: r.Other})
}
