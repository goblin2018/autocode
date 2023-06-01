package house

import (
	"auto/api"
	"auto/app/logic/house"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func update(c *ctx.Context) {
	req := new(api.House)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := house.NewUpdateLogic(c, sv)
	err := l.Update(req)
	c.JSON(nil, err)
}
