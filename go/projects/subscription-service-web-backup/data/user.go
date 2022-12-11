package data

import (
	"context"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

func (u *User) GetAll() ([]*User, error) {
	cxt, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select 
		id,
		email,
		first_name,
		last_name,
		password,
		user_active,
		is_admin,
		created_at,
		updated_at
	from
		users
	order by
		last_name`

	rows, err := db.QueryContext(cxt, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select
		id,
		email,
		first_name,
		last_name,
		password,
		user_active,
		is_admin,
		created_at,
		updated_at
	from
		users
	where
		email = $1`

	var user User
	row := db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select
		id,
		email,
		first_name,
		last_name,
		password,
		user_active,
		is_admin,
		created_at,
		updated_at
	from
		users
	where
		id = $1`

	var user User
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	query = `
	select
		p.id,
		p.plan_name,
		p.plan_amount,
		p.created_at,
		p.updated_at
	from
		user_plans up left join plans p on (p.id = up.plan_id)
	where
		up.user_id = $1`
	var plan Plan
	row = db.QueryRowContext(ctx, query, user.ID)
	err = row.Scan(
		&plan.ID,
		&plan.PlanName,
		&plan.PlanAmount,
		&plan.CreatedAt,
		&plan.UpdatedAt,
	)
	if err == nil {
		user.Plan = &plan
	}
	return &user, nil
}

func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	update
		users
	set
		email = $1,
		first_name = $2,
		last_name = $3,
		user_active = $4,
		updated_at = $5
	where
		id = $6`
	_, err := db.ExecContext(ctx, query,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Active,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	delete from
		users
	where
		id = $1`

	_, err := db.ExecContext(ctx, query, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	delete from
		users
	where
		id = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	query := `
	insert into users
		(email, first_name, last_name, password, user_active, created_at, updated_at
	values
		($1, $2, $3, $4, $5, $6, $7)
	returning id`

	err = db.QueryRowContext(ctx, query,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, nil
	}
	return newID, nil
}

func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	query := `
	update
		users
	set
		password = $1
	where
		id = $2`

	_, err = db.ExecContext(ctx, query, hashedPassword, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) PasswordMatched(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}