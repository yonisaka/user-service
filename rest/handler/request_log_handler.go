package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonisaka/user-service/proto/foo"
)

// RequestLogHandler is a struct
type RequestLogHandler struct {
	*Handler
}

// NewRequestLogHandler is a constructor
func NewRequestLogHandler(h *Handler) *RequestLogHandler {
	return &RequestLogHandler{h}
}

func (r *RequestLogHandler) Create(c *gin.Context) {
	payloads := []*foo.SaveHttpLogRequest{
		{Ip: "1.1.1.1", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.2", Path: "/user/1", Method: "GET"},
		{Ip: "1.1.1.3", Path: "/user/2", Method: "GET"},
		{Ip: "1.1.1.4", Path: "/user/3", Method: "GET"},
		{Ip: "1.1.1.5", Path: "/user", Method: "GET"},
		{Ip: "1.1.1.6", Path: "/user", Method: "GET"},
		{Ip: "1.1.1.7", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.8", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.9", Path: "/user", Method: "POST"},
		{Ip: "1.1.1.10", Path: "/user", Method: "POST"},
	}

	err := r.client.SaveHttpLog(c, payloads)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
