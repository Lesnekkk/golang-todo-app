package users_postgres

import (
	"context"
	"fmt"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (r *UsersPostgresRepository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	query := `
	INSERT INTO todoapp.users (id, version, full_name, phone_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, full_name, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.ID, user.Version, user.FullName, user.PhoneNumber)

	var saved domain.User
	if err := row.Scan(&saved.ID, &saved.Version, &saved.FullName, &saved.PhoneNumber); err != nil {
		return domain.User{}, fmt.Errorf("save user: %w", err)
	}

	return saved, nil
}
