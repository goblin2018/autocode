package sys_user

import (
	"auto/api"
	"auto/app/logic/sys_user"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func sms(c *ctx.Context) {
	req := new(api.SendSmsReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := sys_user.NewSmsLogic(c, sv)
	err := l.Sms(req)
	c.JSON(nil, err)
}
