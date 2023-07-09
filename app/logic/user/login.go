package user

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/user"
	"github.com/jinzhu/copier"
)

type LoginLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewLoginLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *LoginLogic) Login(req *api.User) (err error) {

	return
}
