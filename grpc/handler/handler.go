package handler

import (
	pbAuth "github.com/yonisaka/protobank/auth"
	pbLog "github.com/yonisaka/protobank/log"
	pbUser "github.com/yonisaka/protobank/user"
	"github.com/yonisaka/user-service/config"
	"github.com/yonisaka/user-service/domain/service"
)

// Interface is an interface
type Interface interface {
	pbUser.UserServiceServer
	pbAuth.AuthServer
	pbLog.LogServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	pbUser.UnimplementedUserServiceServer
	pbAuth.UnimplementedAuthServer
	pbLog.UnimplementedLogServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
