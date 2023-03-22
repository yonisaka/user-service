package server

import (
	pbAuth "github.com/yonisaka/protobank/auth"
	pbLog "github.com/yonisaka/protobank/log"
	pbUser "github.com/yonisaka/protobank/user"
	"github.com/yonisaka/user-service/config"
	"github.com/yonisaka/user-service/domain/service"
	"github.com/yonisaka/user-service/grpc/handler"
	"github.com/yonisaka/user-service/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
	repo   *service.Repositories
}

// NewGRPCServer is constructor
func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
	return &Server{
		config: conf,
		repo:   repo,
	}
}

// Run is a method gRPC server
func (s *Server) Run(port int) error {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryLoggerServerInterceptor(),
			interceptor.UnaryAuthServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptor.StreamLoggerServerInterceptor(),
			interceptor.StreamAuthServerInterceptor(),
		),
	)

	handlers := handler.NewHandler(s.config, s.repo)

	// register service server
	pbUser.RegisterUserServiceServer(server, handlers)
	pbAuth.RegisterAuthServer(server, handlers)
	pbLog.RegisterLogServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
