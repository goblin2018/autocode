package api

import (
	"auto/pkg/g"
)

// 系统用户
type SystemUser struct {
	Id       string  `json:"id"`
	Phone    string  `json:"phone" uni:"phone"`
	Name     string  `json:"name"`
	Nickname string  `json:"nickname,omitempty"`
	Duty     string  `json:"duty,omitempty"`
	Avatar   string  `json:"avatar,omitempty"`
	Roles    []int64 `json:"roles"`
}

// 获取阿里云上传凭证
type GetStsResp struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	Expiration      string `json:"expiration"`
}

// 登录
type LoginReq struct {
	Phone   string `json:"phone"`
	Captcha string `json:"captcha"`
}

// 获取系统用户列表
type (
	ListSystemUserReq struct {
		All  bool  `form:"all,omitempty"`
		Page int64 `form:"page,omitempty"`
		Size int64 `form:"size,omitempty"`
	}

	ListSystemUserResp struct {
		Total int64         `json:"total"`
		Items []*SystemUser `json:"items"`
	}
)

// 发送短信验证码
type SendSmsReq struct {
	Phone string `json:"phone"`
}

var systemUserSchema = g.Schema(
	"sys_user",
	g.Ss(SystemUser{}),
	g.Groups(
		&g.G{
			Apis: g.Apis(
				g.Api("update", "", g.PUT, SystemUser{}, g.Empty),
				g.Api("add", "", g.POST, SystemUser{}, g.Empty),
			),
		},

		&g.G{
			Apis: g.Apis(
				g.Api("login", "login", g.POST, LoginReq{}, SystemUser{}),
				g.Api("sms", "sms", g.POST, SendSmsReq{}, g.Empty),
			),
		},
	),

	g.Model(
		"mongo",
		g.Ss(SystemUser{}),

		g.ModelFuncs(),
	),
)
