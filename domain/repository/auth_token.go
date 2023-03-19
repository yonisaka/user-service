package repository

import "github.com/yonisaka/user-service/domain/entity"

type AuthTokenRepository interface {
	CreateAuthToken(*entity.AuthToken, *entity.User) (*entity.AuthToken, error)
	GetAuthToken(int) (*entity.AuthToken, error)
	GetAuthTokenByToken(string) (*entity.AuthToken, error)
}
