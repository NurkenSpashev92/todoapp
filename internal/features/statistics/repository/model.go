package statistics_postgres_repository

import (
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

type StatisticsModel struct {
	TasksCreated                      int
	TasksCompleted                    int
	TasksAverageCompletionTimeSeconds *float64
}

func statisticsDomainFromModel(model StatisticsModel) domain.Statistics {
	var completedRate *float64
	if model.TasksCreated > 0 {
		rate := float64(model.TasksCompleted) / float64(model.TasksCreated) * 100
		completedRate = &rate
	}

	var averageCompletionTime *time.Duration
	if model.TasksAverageCompletionTimeSeconds != nil {
		duration := time.Duration(*model.TasksAverageCompletionTimeSeconds * float64(time.Second))
		averageCompletionTime = &duration
	}

	return domain.NewStatistics(
		model.TasksCreated,
		model.TasksCompleted,
		completedRate,
		averageCompletionTime,
	)
}
