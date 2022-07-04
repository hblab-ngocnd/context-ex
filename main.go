package main

import (
	"context-ex/handlers"
	"context-ex/server"
)

func main() {
	srv := server.DefaultServer()
	srv.Router("aaa", handlers.GetGreeting)
	srv.ListenAndServe()
}
