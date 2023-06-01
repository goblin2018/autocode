package house

import (
	"auto/app/svc"
	"auto/pkg/ctx"
)

var sv *svc.ServiceContext

func RegisterTo(en *ctx.RouterGroup, svc *svc.ServiceContext) {
	sv = svc
	t := en.Group("house")
	{

		t.POST("", add)
		t.PUT("", update)
		t.DELETE("", del)
		t.GET("list", list)
	}
}
