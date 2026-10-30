package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (r *TaskRepository) GetTasks(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed get tasks in repository %w", err)
	}

	var tasksDomain []TaskModel
	for rows.Next() {
		var taskModel TaskModel

		err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorUserId,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		tasksDomain = append(tasksDomain, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	taskDomain := taskDomainsFromModels(tasksDomain)

	return taskDomain, nil
}
