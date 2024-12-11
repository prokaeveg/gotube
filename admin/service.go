package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func FetchAllUsers(db *pgxpool.Pool) ([]ListUser, error) {
	rows, err := db.Query(context.Background(), "SELECT id, username, email, created_at FROM users")

	if err != nil {
		return nil, errors.New("Error fetching all users: " + err.Error())
	}

	defer rows.Close()

	var users []ListUser
	for rows.Next() {
		var user ListUser
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

func CreateUser(db *pgxpool.Pool, user CreateUserRequest) error {
	tx, err := db.Begin(context.Background())

	if err != nil {
		return errors.New("Error starting transaction: " + err.Error())
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		}
	}()

	query := `
	INSERT INTO users (username, email, password_hash, created_at)
	VALUES ($1, $2, $3, $4)
`
	_, err = db.Exec(
		context.Background(),
		query,
		user.Username,
		user.Email,
		user.Password,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return errors.New("Error creating user: " + err.Error())
	}

	return nil
}
