package handler

import (
	"auto/app/handler/house"
	"auto/app/handler/user"
	"auto/app/svc"
	"auto/pkg/ctx"
)

func New(svc *svc.ServiceContext) *ctx.Engine {
	r := ctx.Default()
	api := r.Group("api")
	RegisterRouters(api, svc)
	return r
}

func RegisterRouters(g *ctx.RouterGroup, svc *svc.ServiceContext) {
	user.RegisterTo(g, svc)
	house.RegisterTo(g, svc)

}
