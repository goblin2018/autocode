package user

import (
	"auto/app/svc"
	"auto/pkg/ctx"
)

var sv *svc.ServiceContext

func RegisterTo(en *ctx.RouterGroup, svc *svc.ServiceContext) {
	sv = svc

	t1 := en.Group("user")
	{

		t1.POST("", add)
		t1.POST("", login)
	}

}
