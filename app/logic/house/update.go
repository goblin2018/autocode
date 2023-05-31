package house

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/house"
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

func (l *UpdateLogic) Update(req *api.House) (err error) {

	md := new(house.House)
	copier.Copy(md, req)
	err = l.HouseModel.Update(md)

	return
}
