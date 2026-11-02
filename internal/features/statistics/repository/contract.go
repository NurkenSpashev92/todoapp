package statistics_repository

import (
	"context"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type StatisticsRepositoryInterface interface {
	GetStatistics(ctx context.Context, userId *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}
