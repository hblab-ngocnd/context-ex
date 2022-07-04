package main

import (
	"github.com/hblab-ngocnd/context-ex/handlers"
	"github.com/hblab-ngocnd/context-ex/server"
)

func main() {
	srv := server.DefaultServer()
	srv.Router("aaa", handlers.GetGreeting)
	srv.ListenAndServe()
}
