package users_service

import (
	"context"
	"fmt"

	"github.com/Zakhar4uk/golang-app/internal/core/domain"
)

func (s *USerService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domaun: %w", err)
	}
	user, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}
