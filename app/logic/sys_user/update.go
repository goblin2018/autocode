package sys_user

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/sys_user"
	"github.com/jinzhu/copier"
)

type UpdateLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewUpdateLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *api.SystemUser) (err error) {

	md := new(sys_user.SystemUser)
	copier.Copy(md, req)
	err = l.SysUserModel.Update(md)

	return
}
