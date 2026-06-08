package users_service

import (
	"context"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
)

type USerService struct {
	userRepository UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewUserService(userRepository UserRepository) *USerService {
	return &USerService{
		userRepository: userRepository,
	}
}
