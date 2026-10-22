package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func (r *UserRepository) UpdateUserById(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.users
	SET 
		full_name=$1, phone_number=$2
	WHERE id=$3
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNUmber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user not found by id=%d: %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNUmber,
	)

	return userDomain, nil
}
