package store

import (
	"context"

	"myapp/models"
)

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (s *Pg) CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error) {
	s.logger.Printf("Store. CreateAccount was invoked")

	q := `
INSERT INTO accounts (
	owner,
	balance,
	currency
) VALUES (
	$1, $2, $3
) RETURNING id, owner, balance, currency, created_at
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Account{}, err
	}

	defer stmt.Close()

	// row := s.db.QueryRowContext(ctx, q,
	// 	arg.Owner,
	// 	arg.Balance,
	// 	arg.Currency)

	row := stmt.QueryRowContext(ctx, arg.Owner, arg.Balance, arg.Currency)
	var a models.Account
	err = row.Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)

	return a, err
}

func (s *Pg) GetAccount(ctx context.Context, id int64) (models.Account, error) {
	s.logger.Printf("Store. GetAccount was invoked")

	q := `
SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Account{}, err
	}

	defer stmt.Close()

	var a models.Account
	err = stmt.QueryRowContext(ctx, id).Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)

	// row := s.db.QueryRowContext(ctx, q, id)

	// var a models.Account
	// err := row.Scan(
	// 	&a.ID,
	// 	&a.Owner,
	// 	&a.Balance,
	// 	&a.Currency,
	// 	&a.CreatedAt,
	// )
	return a, err
}

func (s *Pg) GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error) {
	s.logger.Printf("Store. GetAccountForUpdate was invoked")

	q := `
SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 FOR NO KEY UPDATE
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Account{}, err
	}

	defer stmt.Close()

	var a models.Account
	err = stmt.QueryRowContext(ctx, id).Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)

	// row := s.db.QueryRowContext(ctx, q, id)

	// var a models.Account
	// err := row.Scan(
	// 	&a.ID,
	// 	&a.Owner,
	// 	&a.Balance,
	// 	&a.Currency,
	// 	&a.CreatedAt,
	// )
	return a, err
}

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (s *Pg) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]models.Account, error) {
	s.logger.Printf("Store. ListAccounts was invoked")

	q := `
SELECT id, owner, balance, currency, created_at FROM accounts
	ORDER BY id
	LIMIT $1
	OFFSET $2
`

	rows, err := s.db.QueryContext(ctx, q, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]models.Account, 0)
	for rows.Next() {
		var a models.Account
		if err = rows.Scan(
			&a.ID,
			&a.Owner,
			&a.Balance,
			&a.Currency,
			&a.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, a)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (s *Pg) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (models.Account, error) {
	s.logger.Printf("Store. UpdateAccount was invoked")

	q := `
UPDATE accounts
	SET balance = $2
	WHERE id = $1
	RETURNING id, owner, balance, currency, created_at
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Account{}, err
	}

	defer stmt.Close()

	var a models.Account
	err = stmt.QueryRowContext(ctx, arg.ID, arg.Balance).Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)

	// _, err := s.db.ExecContext(ctx, q, arg.ID, arg.Balance)
	return a, err
}

func (s *Pg) DeleteAccount(ctx context.Context, id int64) error {
	s.logger.Printf("Store. DeleteAccount was invoked")

	q := `
DELETE FROM accounts
	WHERE id = $1
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)

	// _, err := s.db.ExecContext(ctx, q, id)
	return err
}
