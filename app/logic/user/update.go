package user

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/user"
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

func (l *UpdateLogic) Update(req *api.User) (err error) {

	md := new(user.User)
	copier.Copy(md, req)
	err = l.UserModel.Update(md)

	return
}
