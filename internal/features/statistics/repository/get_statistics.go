package statistics_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (r *StatisticsRepository) GetStatistics(
	ctx context.Context,
	userId *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT
		COUNT(*),
		COUNT(*) FILTER (WHERE completed),
		EXTRACT(EPOCH FROM AVG(completed_at - created_at) FILTER (WHERE completed))::FLOAT8
	FROM todoapp.tasks
	WHERE ($1::INTEGER IS NULL OR author_user_id=$1)
		AND ($2::TIMESTAMPTZ IS NULL OR created_at >= $2)
		AND ($3::TIMESTAMPTZ IS NULL OR created_at < $3);
	`

	row := r.pool.QueryRow(ctx, query, userId, from, to)

	var statisticsModel StatisticsModel

	err := row.Scan(
		&statisticsModel.TasksCreated,
		&statisticsModel.TasksCompleted,
		&statisticsModel.TasksAverageCompletionTimeSeconds,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("failed to scan statistics: %w", err)
	}

	return statisticsDomainFromModel(statisticsModel), nil
}
