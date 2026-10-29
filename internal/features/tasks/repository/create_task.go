package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
)

func (r *TaskRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks(title, description, completed, created_at, completed_at, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserId,
	)

	var taskModel TaskModel
	err := row.Scan(
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
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserId,
	)

	return taskDomain, nil
}
