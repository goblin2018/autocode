package main

import (
	"auto/api"
	"auto/app/config"

	"auto/pkg/g"
	"auto/pkg/log"
)

func main() {
	c := config.New()
	log.Init(c.Mode, c.LogLevel)

	schemas := api.Schemas
	baseDir := api.BaseDir

	for _, schema := range schemas {
		g.GenTypes(baseDir, schema.Name, schema.Structs)
		g.GenModels(baseDir, schema.Name, schema.Model)
		g.GenLogics(baseDir, schema.Name, schema)
		g.GenHandlers(baseDir, schema)
	}
}
