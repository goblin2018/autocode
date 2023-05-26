package api

import (
	"auto/pkg/generator/g"
)

// db: mongo
type User struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var get = g.Api("Get", User{}, g.Empty)
var update = g.Api("Update", User{}, g.Empty)

var userSchema = g.Schema(
	g.Structs(User{}),
	g.Groups(
		g.Group("user",
			g.Apis(get, update),
		),
	),
)
