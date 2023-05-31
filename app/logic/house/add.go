package house

import (
	"auto/api"
	"auto/app/model/house"
	"auto/app/svc"
	"auto/pkg/log"
	"context"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type AddLogic struct {
	*zap.SugaredLogger
	ctx context.Context
	*svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		SugaredLogger:  log.L,
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
