package repository

import (
	"authService/internal/domain"
	"context"
	"database/sql"
	"time"
)

type user struct {
	conn *sql.DB
}

type User interface {
	SaveUser(ctx context.Context, user domain.User) error
}

func New(conn *sql.DB) User {
	return &user{
		conn: conn,
	}
}

func (u user) SaveUser(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := u.conn.ExecContext(ctx,
		"INSERT INTO users(name, username, password) VALUES($1, $2, $3)",
		user.Name,
		user.Username,
		user.Password,
	); err != nil {
		return err
	}

	return nil
}
