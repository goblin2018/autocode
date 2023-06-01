package user

import (
	"auto/app/svc"
	"auto/pkg/ctx"
)

var sv *svc.ServiceContext

func RegisterTo(en *ctx.RouterGroup, svc *svc.ServiceContext) {
	sv = svc
	t := en.Group("user")
	{

		t.PUT("", update)
		t.POST("", add)
	}
}
