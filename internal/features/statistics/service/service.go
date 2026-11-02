package statistics_service

import (
	statistics_repository "github.com/nurkenspashev92/todoapp/internal/features/statistics/repository"
)

type StatisticsService struct {
	statisticsRepository statistics_repository.StatisticsRepositoryInterface
}

func NewStatisticsService(statisticsRepository statistics_repository.StatisticsRepositoryInterface) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}
