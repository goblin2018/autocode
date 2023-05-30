package api

import "auto/pkg/g"

var Schemas []g.Sc
var BaseDir = "app"

func Add(s ...g.Sc) {
	Schemas = append(Schemas, s...)
}

func init() {
	Add(
		userSchema,
		houseSchema,
	)
}
