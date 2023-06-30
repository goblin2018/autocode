package house

import (
	"auto/app/svc"
	"auto/pkg/ctx"
)

var sv *svc.ServiceContext

func RegisterTo(en *ctx.RouterGroup, svc *svc.ServiceContext) {
	sv = svc

	t1 := en.Group("house")
	{

		t1.POST("", add)
		t1.PUT("", update)
		t1.DELETE("", del)
		t1.GET("list", list)
	}

}
