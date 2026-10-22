package users_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (s *UserSerivce) PatchUserById(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := s.userRepository.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("user not found: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := s.userRepository.PatchUserById(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("faild to update user from repository: %w", err)
	}

	return patchedUser, nil
}
