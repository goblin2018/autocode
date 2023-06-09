package house

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/house"
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

func (l *AddLogic) Add(req *api.House) (err error) {

	md := new(house.House)
	copier.Copy(md, req)
	err = l.HouseModel.Create(md)

	return
}
