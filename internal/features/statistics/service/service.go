package statistics_service

import (
	"context"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepositoryInterface
}

type StatisticsRepositoryInterface interface {
	GetStatistics(ctx context.Context, userId *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsService(statisticsRepository StatisticsRepositoryInterface) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}
