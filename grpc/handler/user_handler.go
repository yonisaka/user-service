package handler

import (
	"context"

	pb "github.com/yonisaka/protobank/user"
	"github.com/yonisaka/user-service/domain/entity"
	"github.com/yonisaka/user-service/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser is a function
func (c *Handler) GetUser(ctx context.Context, r *pb.UserByIDRequest) (*pb.UserResponse, error) {
	usr, err := c.repo.User.Find(ctx, int(r.GetId()))

	if err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	return &pb.UserResponse{
		Id:        uint64(usr.ID),
		Name:      usr.Name,
		Username:  usr.Username,
		CreatedAt: usr.CreatedAt.String(),
		UpdatedAt: usr.UpdatedAt.String(),
	}, nil
}

// UpdateUser is function
func (c *Handler) UpdateUser(ctx context.Context, payload *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	userId := int(payload.GetId())

	if _, err := c.repo.User.Find(ctx, userId); err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	userData := &entity.User{
		Name:     payload.GetName(),
		Username: payload.GetUsername(),
	}

	err := c.repo.User.Update(ctx, userId, userData)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(userData.ID),
		Name:      userData.Name,
		Username:  userData.Username,
		CreatedAt: userData.CreatedAt.String(),
		UpdatedAt: userData.UpdatedAt.String(),
	}, nil
}

// CreateUser is function
func (c *Handler) CreateUser(ctx context.Context, r *pb.UserCreateRequest) (*pb.UserResponse, error) {
	usr := entity.User{
		Name:     r.GetName(),
		Username: r.GetUsername(),
		Password: utils.Hmac256(r.GetPassword(), utils.HmacSecret()),
	}

	err := c.repo.User.Create(ctx, &usr)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(usr.ID),
		Name:      usr.Name,
		Username:  usr.Username,
		CreatedAt: usr.CreatedAt.String(),
		UpdatedAt: usr.UpdatedAt.String(),
	}, nil
}

// GetUserList is function
func (c *Handler) GetUserList(ctx context.Context, _ *pb.UserListQuery) (*pb.UsersResponse, error) {
	serv, err := c.repo.User.Get(ctx)

	if err != nil {
		return nil, err
	}

	var users []*pb.UserResponse
	for _, u := range serv {
		users = append(users, &pb.UserResponse{
			Id:        uint64(u.ID),
			Name:      u.Name,
			Username:  u.Username,
			CreatedAt: u.CreatedAt.String(),
			UpdatedAt: u.UpdatedAt.String(),
		})
	}

	return &pb.UsersResponse{
		Users: users,
	}, nil
}

// DeleteUser is a function
func (c *Handler) DeleteUser(ctx context.Context, r *pb.UserByIDRequest) (*pb.UserDeleteResponse, error) {
	userId := int(r.GetId())

	if _, err := c.repo.User.Find(ctx, userId); err != nil {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	err := c.repo.User.Delete(ctx, int(r.GetId()))

	if err != nil {
		return nil, err
	}

	return &pb.UserDeleteResponse{
		Message: "ok",
	}, nil
}
