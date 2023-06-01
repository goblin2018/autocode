package user

import (
	"auto/api"
	"auto/app/logic/user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func update(c *ctx.Context) {
	req := new(api.User)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := user.NewUpdateLogic(c, sv)
	err := l.Update(req)
	c.JSON(nil, err)
}
