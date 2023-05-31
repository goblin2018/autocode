package handler

import (
	"auto/app/handler/house"
	"auto/pkg/ctx"
)

func New() *ctx.Engine {
	r := ctx.Default()
	api := r.Group("/api")
	house.Register(api)
	return r
}
