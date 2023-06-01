package house

import (
	"auto/api"
	"auto/app/logic/house"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func list(c *ctx.Context) {
	req := new(api.ListHouseReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := house.NewListLogic(c, sv)
	resp, err := l.List(req)
	c.JSON(resp, err)
}
