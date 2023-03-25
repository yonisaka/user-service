package handler

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	pb "github.com/yonisaka/protobank/auth"
	"github.com/yonisaka/user-service/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthLogin is a function
func (c *Handler) AuthLogin(ctx context.Context, r *pb.AuthLoginPayload) (*pb.LoginResponse, error) {
	username := r.GetUsername()
	password := r.GetPassword()

	us, err := c.repo.User.FindByUsername(ctx, username)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid username or password")
	}

	if !utils.HmacComparator(password, us.Password, utils.HmacSecret()) {
		return nil, status.Error(codes.InvalidArgument, "Invalid username or password")
	}

	return &pb.LoginResponse{
		Ok:          true,
		AccessToken: utils.EncodeBasicAuth(us.Username, password),
	}, nil
}

func (c *Handler) AuthB2B(ctx context.Context, r *pb.AuthB2BPayload) (*pb.UserResponse, error) {
	token := r.GetToken()

	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "no token provided")
	}

	claims := &utils.JwtClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.config.JWT.SignatureKey), nil
	})

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token provided")
	}

	if !jwtToken.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token provided")
	}

	us, err := c.repo.User.FindByUsername(ctx, claims.Username)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid username or password")
	}

	return &pb.UserResponse{
		Id:       uint64(us.ID),
		Name:     us.Name,
		Username: us.Username,
	}, nil
}
