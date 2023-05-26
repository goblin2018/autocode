package main

import (
	"auto/app/handler"
	"auto/pkg/server"

	"auto/app/config"
)

func main() {
	c := config.New()
	r := handler.New()
	srv := server.New(c.Port, r)
	srv.Start()
	defer srv.Stop()
}
