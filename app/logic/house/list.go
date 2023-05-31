package house

import (
	"auto/api"
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/log"

	"auto/app/model/house"
	"github.com/jinzhu/copier"
)

type ListLogic struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func NewListLogic(ctx *ctx.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *ListLogic) List(req *api.ListHouseReq) (resp *api.ListHouseResp, err error) {

	resp = new(api.ListHouseResp)
	opt := new(house.ListHouseReq)
	copier.Copy(opt, req)
	items, _ := l.HouseModel.List(l.ctx, opt)
	for _, it := range items {
		td := new(api.House)
		copier.Copy(td, it)
		resp.Items = append(resp.Items, td)
	}

	if opt.Size != 0 {
		resp.Total, _ = l.HouseModel.Count(l.ctx, opt)
	}

	return
}
