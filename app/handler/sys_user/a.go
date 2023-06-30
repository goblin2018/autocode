package sys_user

import (
	"auto/app/svc"
	"auto/pkg/ctx"
)

var sv *svc.ServiceContext

func RegisterTo(en *ctx.RouterGroup, svc *svc.ServiceContext) {
	sv = svc

	t1 := en.Group("sys_user")
	{

		t1.PUT("", update)
		t1.POST("", add)
	}

	t2 := en.Group("sys_user")
	{

		t2.POST("login", login)
		t2.POST("sms", sms)
	}

}
