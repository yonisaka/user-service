package handler

import (
	"context"
	"fmt"

	"github.com/yonisaka/user-service/proto/foo"
)

// SayHello is function
func (c *Handler) SayHello(_ context.Context, r *foo.HelloRequest) (*foo.HelloReply, error) {
	return &foo.HelloReply{Message: fmt.Sprintf("Hello: %v", r.GetText())}, nil
}

// Ping is a function
func (c *Handler) Ping(_ context.Context, _ *foo.PingRequest) (*foo.PingReply, error) {
	return &foo.PingReply{Redis: "Ok", Db: "Ok"}, nil
}
