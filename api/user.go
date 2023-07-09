package api

import (
	"auto/pkg/g"
)

type User struct {
	Id        string `json:"id"`
	OpenId    string `json:"openId" uni:"openId"` // 微信openId
	Avatar    string `json:"avatar"`              // 头像
	Phone     string `json:"phone" uni:"phone"`   // 手机号，唯一
	Name      string `json:"name"`                // 真实姓名
	Nickname  string `json:"nickname"`            // 昵称
	AddressId string `json:"addressId"`           // 收货地址Id, 用于加载收获地址
	Birthday  int64  `json:"birthday"`            // 生日，时间戳
	State     int64  `json:"state"`               // 状态，1正常，2禁用
	ExpireAt  int64  `json:"expireAt"`            // 会员到期时间，时间戳
}

var userSchema = g.Schema(
	"user",
	g.Ss(User{}),
	g.Groups(
		&g.G{
			Apis: g.Apis(
				&g.A{
					Name:   "add",
					Path:   "",
					Method: g.POST,
					Input:  User{},
					Output: g.Empty,
				},
				&g.A{
					Name:   "login",
					Path:   "",
					Method: g.POST,
					Input:  User{},
					Output: g.Empty,
				},
			),
		},
	),

	g.Model(
		"mysql",
		g.Ss(User{}),
		g.ModelFuncs(),
	),
)
