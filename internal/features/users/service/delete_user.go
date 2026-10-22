package users_service

import (
	"context"
	"fmt"
)

func (s *UserSerivce) DeleteUserById(ctx context.Context, userId int) error {
	if err := s.userRepository.DeleteUserById(ctx, userId); err != nil {
		return fmt.Errorf("faild to delete user from repository: %w", err)
	}

	return nil
}
