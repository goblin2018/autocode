package {{.package}}

import (
	"auto/api"
	"auto/pkg/log"
	"auto/pkg/ctx"
	"auto/app/svc"
	{{.imports}}
)

type {{.logic}} struct {
	*log.Logger
	ctx *ctx.Context
	*svc.ServiceContext
}

func New{{.logic}}(ctx *ctx.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger:         log.L,
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) ({{.response}}err error) {
	{{.body}}
	return
}
