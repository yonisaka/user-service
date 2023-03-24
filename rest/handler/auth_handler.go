package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	ExpiredAt   string `json:"expired_at"`
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

	expirationTime := time.Now().Add(time.Duration(r.config.JWT.Expiration) * time.Minute)
	claims := &utils.JwtClaims{
		ID:       int(us.ID),
		Username: us.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(r.config.JWT.SignatureKey))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage("Error while signing token"),
		)
		return
	}
	res.AccessToken = tokenString
	res.ExpiredAt = expirationTime.Format(time.RFC1123)
	c.JSON(http.StatusOK, *dto.NewResponse().WithCode(http.StatusOK).WithData(res))
}
