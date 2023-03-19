package route

import (
	"github.com/yonisaka/user-service/config"
	"github.com/yonisaka/user-service/domain/service"
	"github.com/yonisaka/user-service/grpc/client"
)

// Router is a struct contains dependencies needed
type Router struct {
	config *config.Config
	repo   *service.Repositories
	client *client.GRPCClient
}

// RouterOption return Router with RouterOption to fill up the dependencies
type RouterOption func(*Router)

// NewRouter is a constructor will initialize Router.
func NewRouter(options ...RouterOption) *Router {
	router := &Router{}

	for _, opt := range options {
		opt(router)
	}

	return router
}
