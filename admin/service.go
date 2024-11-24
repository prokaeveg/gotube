package admin

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func FetchAllUsers(db *pgxpool.Pool) ([]User, error) {
	rows, err := db.Query(context.Background(), "SELECT id, username, email, created_at FROM users")

	if err != nil {
		return nil, errors.New("Error fetching all users: " + err.Error())
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, errors.New("Error scanning user: " + err.Error())
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("Error reading rows: " + err.Error())
	}

	return users, nil
}
