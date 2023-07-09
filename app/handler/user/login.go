package user

import (
	"auto/api"
	"auto/app/logic/user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func login(c *ctx.Context) {
	req := new(api.User)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := user.NewLoginLogic(c, sv)
	err := l.Login(req)
	c.JSON(nil, err)
}
