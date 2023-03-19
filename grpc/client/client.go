package client

import (
	"github.com/yonisaka/user-service/proto/foo"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	httpLog foo.LogServiceClient
	hello   foo.HelloClient
}

// NewGRPCClient is constructor
func NewGRPCClient(conn grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		httpLog: foo.NewLogServiceClient(conn),
		hello:   foo.NewHelloClient(conn),
	}
}
