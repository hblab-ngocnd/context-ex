package main

import (
	"github.com/hblab-ngocnd/context-ex/server"

	"github.com/hblab-ngocnd/context-ex/handlers"
)

func main() {
	srv := server.DefaultServer()
	srv.Router("aaa", handlers.GetGreeting)
	srv.ListenAndServe()
}
