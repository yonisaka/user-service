package service

import (
	"github.com/yonisaka/user-service/domain/repository"
	"github.com/yonisaka/user-service/infrastructure/persistence"
	"gorm.io/gorm"
)

// Repositories is a struct
type Repositories struct {
	User      repository.UserRepositoryInterface
	AuthToken repository.AuthTokenRepository
	HttpLog   repository.HttpLogRepositoryInterface
	DB        *gorm.DB
}

// NewDBService is constructor
func NewDBService(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    persistence.NewUserRepository(db),
		HttpLog: persistence.NewHttpLogRepository(db),
		DB:      db,
	}
}
