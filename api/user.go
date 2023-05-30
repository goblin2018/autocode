package api

import (
	"auto/pkg/g"
)

type User struct {
	Id    string `json:"id"`
	Phone string `json:"phone" uni:"phone"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
}

type TestB struct {
	LaiLe string `json:"name"`
	Age   int64  `json:"age"`
}

var get = g.Api("get", User{}, g.Empty)
var update = g.Api("Update", User{}, g.Empty)

var userSchema = g.Schema(
	"user",
	g.Ss(User{}, TestB{}),
	g.Groups(
		g.Group("user",
			g.Apis(get, update),
		),
	),

	g.Model(
		"mongo",
		g.Ss(User{}),

		g.ModelFuncs(
			g.ModelFunc("get", User{}, User{}),
			g.ModelFunc("update", User{}, User{}),
		),
	),
)
