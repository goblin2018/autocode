package user

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/user"
	"github.com/jinzhu/copier"
)

type AddLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewAddLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *AddLogic) Add(req *api.User) (err error) {

	md := new(user.User)
	copier.Copy(md, req)
	err = l.UserModel.Create(md)

	return
}
