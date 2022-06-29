package server

import (
	"context"
	"net"

	"github.com/hblab-ngocnd/context-ex/session"

	"github.com/hblab-ngocnd/context-ex/handlers"
)

type MyServer struct {
	router map[string]handlers.MyHandlerFunc
}

func (sv *MyServer) Router(url string, handlerFunc handlers.MyHandlerFunc) {
	if sv.router == nil {
		sv.router = make(map[string]handlers.MyHandlerFunc)
	}
	sv.router[url] = handlerFunc
}
func (sv *MyServer) ListenAndServe() {
	addr, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		ctx := session.SetSessionID(context.Background())
		go sv.Request(ctx, conn)
	}
}
func DefaultServer() *MyServer {
	return &MyServer{}
}
func (sv *MyServer) Request(ctx context.Context, conn *net.TCPConn) {
	req := handlers.MyRequest{}
	req.SetConn(conn)
	if handler, ok := sv.router[req.GetPath()]; ok {
		handler(ctx, req)
	} else {
		handlers.NotFoundHandler(ctx, req)
	}
}
