package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/user-service/config"
	"github.com/yonisaka/user-service/domain/service"
	"github.com/yonisaka/user-service/grpc/client"
	"github.com/yonisaka/user-service/rest/handler"
	"github.com/yonisaka/user-service/rest/middleware"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(gClient *client.GRPCClient) RouterOption {
	return func(r *Router) {
		r.client = gClient
	}
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(middleware.Logger())

	hand := handler.NewHandler(r.repo, r.client)

	httpLog := handler.NewRequestLogHandler(hand)
	hello := handler.NewHelloHandler(hand)

	e.GET("/api/ping", hello.Ping)
	e.POST("/api/hello", hello.SayHello)
	e.POST("/api/request-logs", httpLog.Create)

	return e
}
