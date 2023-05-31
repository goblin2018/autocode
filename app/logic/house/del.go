package house

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"
)

type DelLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewDelLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *DelLogic) Del(req *api.DelHouseReq) (err error) {

	err = l.HouseModel.Delete(req.Id)

	return
}
