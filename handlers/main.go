package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"context-ex/auth"
	"context-ex/db"
)

type MyHandlerFunc func(context.Context, MyRequest)

type MyResponse struct {
	Code int
	Body string
	Err  error
}
type MyRequest struct {
	path string
	conn net.Conn
}

func (r *MyRequest) SetPath(path string) {
	r.path = path
}
func (r *MyRequest) GetPath() string {
	return r.path
}
func (r *MyRequest) SetConn(conn net.Conn) {
	r.conn = conn
}

var GetGreeting MyHandlerFunc = func(ctx context.Context, req MyRequest) {
	var res MyResponse
	// トークンからユーザー検証→ダメなら即return
	userID, err := auth.VerifyAuthToken(ctx)
	if err != nil {
		res = MyResponse{Code: 403, Err: err}
		req.Send(res)
		return
	}
	// DBリクエストをいつタイムアウトさせるかcontext経由で設定
	dbReqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)

	//DBからデータ取得
	rcvChan := db.DefaultDB.Search(dbReqCtx, userID)
	data, ok := <-rcvChan
	cancel()

	// DBリクエストがタイムアウトしていたら408で返す
	if !ok {
		res = MyResponse{Code: 408, Err: errors.New("DB request timeout")}
		req.Send(res)
		return
	}
	// レスポンスの作成
	res = MyResponse{
		Code: 200,
		Body: fmt.Sprintf("From path %s, Hello! your ID is %d\ndata → %s", req.path, userID, data),
	}

	// レスポンス内容を標準出力(=本物ならnet.Conn)に書き込み
	req.Send(res)
}

var NotFoundHandler MyHandlerFunc = func(ctx context.Context, req MyRequest) {
	var res MyResponse
	res = MyResponse{
		Code: 404,
		Body: fmt.Sprintf("Not found"),
	}
	req.Send(res)
}

func (r MyRequest) Send(data MyResponse) {
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	r.conn.Write([]byte("HTTP/1.1  200 OK\r\n"))
	r.conn.Write([]byte("Content-Type: application/json\r\n"))
	r.conn.Write([]byte(fmt.Sprintf("Content-Length: %d\r\n", len(body))))
	r.conn.Write([]byte("\r\n"))
	r.conn.Write(body)
	r.conn.Close()
	return
}
