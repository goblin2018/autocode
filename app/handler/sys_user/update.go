package sys_user

import (
	"auto/api"
	"auto/app/logic/sys_user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func update(c *ctx.Context) {
	req := new(api.SystemUser)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := sys_user.NewUpdateLogic(c, sv)
	err := l.Update(req)
	c.JSON(nil, err)
}
