package sys_user

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/sys_user"
	"github.com/jinzhu/copier"
)

type SmsLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewSmsLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *SmsLogic {
	return &SmsLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *SmsLogic) Sms(req *api.SendSmsReq) (err error) {

	return
}
