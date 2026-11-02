package users_service

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func (s *UserSerivce) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvlidArgument,
		)
	}
	if limit != nil && *limit > domain.MaxLimit {
		return nil, fmt.Errorf(
			"limit must be <= %d: %w",
			domain.MaxLimit,
			core_errors.ErrInvlidArgument,
		)
	}
	if limit == nil {
		defaultLimit := domain.DefaultLimit
		limit = &defaultLimit
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvlidArgument,
		)
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("faild to get users: %w", err)
	}

	return users, nil
}
