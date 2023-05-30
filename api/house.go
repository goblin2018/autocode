package api

import (
	"auto/pkg/g"
)

type House struct {
	Id    string  `json:"id"`
	OrgId string  `json:"orgId" uni:"org_phone"`
	Phone string  `json:"phone" uni:"org_phone"`
	Name  string  `json:"name"`
	Addr  string  `json:"addr"`
	Users []*User `json:"users" api:"true"`
}

type (
	ListHouseReq struct {
		OrgId string `form:"orgId,omitempty"`
		Page  int64  `form:"page,omitempty"`
		Size  int64  `form:"size,omitempty"`
	}

	ListHouseResp struct {
		Items []*House `json:"items" load:"Users"`
		Total int64    `json:"total"`
	}
)

type DelHouseReq struct {
	Id string `json:"id"`
}

var houseSchema = g.Schema(
	"house",
	g.Ss(House{}, ListHouseReq{}, ListHouseResp{}, DelHouseReq{}),
	g.Groups(
		g.Group("house",
			g.Apis(
				g.Api("add", House{}, g.Empty),
				g.Api("update", House{}, g.Empty),
				g.Api("del", DelHouseReq{}, g.Empty),
				g.Api("list", ListHouseReq{}, ListHouseResp{}),
			),
		),
	),

	g.Model(
		"mongo",
		g.Ss(House{}),
		g.ModelFuncs(
			g.ModelFunc("get", User{}, User{}),
			g.ModelFunc("update", User{}, User{}),
		),
	),
)