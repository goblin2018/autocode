package main

import (
	"auto/app/handler"
	"auto/app/svc"
	"auto/pkg/log"
	"auto/pkg/server"

	"auto/app/config"
)

func main() {
	c := config.New()
	log.Init(c.Mode, c.LogLevel)
	log.Info("init logger")
	svc := svc.NewServiceContext(c)
	log.Info("init service context")

	r := handler.New(svc)
	srv := server.New(c.Port, r)
	srv.Start()
	defer srv.Stop()
}
