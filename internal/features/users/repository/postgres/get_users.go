package users_postgres

import (
	"context"
	"fmt"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (r *UsersPostgresRepository) GetMany(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	ORDER BY id ASC
	LIMIT $1 OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get many users: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Version, &user.FullName, &user.PhoneNumber); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}
