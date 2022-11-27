package store

import (
	"context"
	"myapp/models"
)

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (s *Pg) CreateUser(ctx context.Context, arg CreateUserParams) (models.User, error) {
	s.logger.Printf("Store. CreateUser was invoked")

	q := `
INSERT INTO users (
	username,
	hashed_password,
	full_name,
	email
) VALUES (
	$1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, arg.Username, arg.HashedPassword, arg.FullName, arg.Email)
	var u models.User
	err = row.Scan(
		&u.Username,
		&u.HashedPassword,
		&u.FullName,
		&u.Email,
		&u.PasswordChangedAt,
		&u.CreatedAt,
	)

	return u, err
}

func (s *Pg) GetUser(ctx context.Context, username string) (models.User, error) {
	s.logger.Printf("Store. GetUser was invoked")

	q := `
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
	WHERE username = $1 LIMIT 1
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, username)
	var u models.User
	err = row.Scan(
		&u.Username,
		&u.HashedPassword,
		&u.FullName,
		&u.Email,
		&u.PasswordChangedAt,
		&u.CreatedAt,
	)

	return u, err
}
