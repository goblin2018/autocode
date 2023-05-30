package main

import (
	"auto/app/handler"
	"auto/pkg/log"
	"auto/pkg/server"

	"auto/app/config"
)

func main() {
	c := config.New()
	log.Init(c.Mode, c.LogLevel)
	r := handler.New()
	srv := server.New(c.Port, r)
	srv.Start()
	defer srv.Stop()
}
