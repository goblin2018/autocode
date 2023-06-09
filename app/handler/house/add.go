package house

import (
	"auto/api"
	"auto/app/logic/house"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func add(c *ctx.Context) {
	req := new(api.House)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := house.NewAddLogic(c, sv)
	err := l.Add(req)
	c.JSON(nil, err)
}
