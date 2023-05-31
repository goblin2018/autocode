package house

import "auto/pkg/ctx"

func Register(en *ctx.RouterGroup) {
	t := en.Group("/house")

	{
		t.GET("", add)
	}
}
