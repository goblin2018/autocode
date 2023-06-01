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

var userSchema = g.Schema(
	"user",
	g.Ss(User{}, TestB{}),
	g.Groups(
		g.Group("user",
			g.Api("update", "", User{}, g.Empty),
			g.Api("add", "", User{}, g.Empty),
		),
	),

	g.Model(
		"mongo",
		g.Ss(User{}),

		g.ModelFuncs(),
	),
)
