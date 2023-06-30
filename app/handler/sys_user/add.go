package sys_user

import (
	"auto/api"
	"auto/app/logic/sys_user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func add(c *ctx.Context) {
	req := new(api.SystemUser)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := sys_user.NewAddLogic(c, sv)
	err := l.Add(req)
	c.JSON(nil, err)
}
