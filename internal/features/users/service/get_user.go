package users_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *UserSerivce) GetUserById(ctx context.Context, userId int) (domain.User, error) {
	user, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("faild to get user from repository: %w", err)
	}

	return user, nil
}
