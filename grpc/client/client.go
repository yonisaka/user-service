package client

import (
	pbLog "github.com/yonisaka/protobank/log"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	httpLog pbLog.LogServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(conn grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		httpLog: pbLog.NewLogServiceClient(conn),
	}
}
