package sys_user

import (
	"auto/api"
	"auto/app/logic/sys_user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func login(c *ctx.Context) {
	req := new(api.LoginReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := sys_user.NewLoginLogic(c, sv)
	resp, err := l.Login(req)
	c.JSON(resp, err)
}
