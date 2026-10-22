package users_service

import (
	"context"
	"fmt"
)

func (s *UserSerivce) DeleteUserById(ctx context.Context, id int) error {
	if err := s.userRepository.DeleteUserById(ctx, id); err != nil {
		return fmt.Errorf("faild to delete user from repository: %w", err)
	}

	return nil
}
