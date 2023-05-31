package house

import (
	"auto/api"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func add(c *ctx.Context) {
	req := new(api.House)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

}
