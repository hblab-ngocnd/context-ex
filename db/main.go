package db

import "context"

type db struct {
}

var DefaultDB = func() *db {
	return &db{}
}()

func (d *db) Search(ctx context.Context, userID int) <-chan interface{} {
	out := make(chan interface{})
	select {
	case <-ctx.Done():
	}
	return out
}
