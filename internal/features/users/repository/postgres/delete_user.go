package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func (r *UserRepository) DeleteUserById(ctx context.Context, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.users
	WHERE id=$1
	`

	cmdTag, err := r.pool.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d' not found: %w", userId, core_errors.ErrNotFound)
	}

	return nil
}
