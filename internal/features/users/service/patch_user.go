package users_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *UserSerivce) UpdateUserById(ctx context.Context, id int, user domain.User) (domain.User, error) {
	user, err := s.userRepository.UpdateUserById(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("faild to get user from repository: %w", err)
	}

	return user, nil
}
