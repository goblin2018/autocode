package handler

import "auto/pkg/ctx"

func New() *ctx.Engine {
	r := ctx.Default()
	return r
}
