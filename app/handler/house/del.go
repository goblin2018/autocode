package house

import (
	"auto/api"
	"auto/app/logic/house"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func del(c *ctx.Context) {
	req := new(api.DelHouseReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := house.NewDelLogic(c, sv)
	err := l.Del(req)
	c.JSON(nil, err)
}
