package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userId *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if userId != nil && *userId <= 0 {
		return domain.Statistics{}, fmt.Errorf(
			"user_id must be positive: %w",
			core_errors.ErrInvlidArgument,
		)
	}

	if from != nil && to != nil && !from.Before(*to) {
		return domain.Statistics{}, fmt.Errorf(
			"`from` must be before `to`: %w",
			core_errors.ErrInvlidArgument,
		)
	}

	statistics, err := s.statisticsRepository.GetStatistics(ctx, userId, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("faild to get statistics from repository: %w", err)
	}

	return statistics, nil
}
