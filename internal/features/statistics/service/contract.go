package statistics_service

import (
	"context"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

var _ StatisticsServiceInterface = (*StatisticsService)(nil)

type StatisticsServiceInterface interface {
	GetStatistics(ctx context.Context, userId *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}
