package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonisaka/user-service/rest/dto"
	"github.com/yonisaka/user-service/utils"
)

type AuthHandler struct {
	*Handler
}

func NewAuthHandler(h *Handler) *AuthHandler {
	return &AuthHandler{h}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (r *AuthHandler) AuthLogin(c *gin.Context) {
	var (
		req LoginRequest
		res LoginResponse
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	us, err := r.repo.User.FindByUsername(c, req.Username)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusUnauthorized).WithMessage("Invalid username or password"),
		)
		return
	}

	if !utils.HmacComparator(req.Password, us.Password, utils.HmacSecret()) {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusUnauthorized).WithMessage("Invalid username or password"),
		)
		return
	}

	res.AccessToken = utils.EncodeBasicAuth(us.Username, us.Password)
	c.JSON(http.StatusOK, *dto.NewResponse().WithCode(http.StatusOK).WithData(res))
}
