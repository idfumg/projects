package store

import (
	"context"
	"myapp/models"
)

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (s *Pg) CreateEntry(ctx context.Context, arg CreateEntryParams) (models.Entry, error) {
	s.logger.Printf("Store. CreateEntry was invoked")

	q := `
INSERT INTO entries (
	account_id,
	amount
) VALUES (
	$1, $2
) RETURNING id, account_id, amount, created_at
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		s.logger.Printf("PrepareContext error")
		return models.Entry{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, arg.AccountID, arg.Amount)
	var e models.Entry
	err = row.Scan(&e.ID, &e.AccountID, &e.Amount, &e.CreatedAt)
	if err != nil {
		s.logger.Printf("Scan error")
	}

	// row := s.db.QueryRowContext(ctx, q, arg.AccountID, arg.Amount)

	// var e models.Entry
	// err := row.Scan(&e.ID, &e.AccountID, &e.Amount, &e.CreatedAt)

	return e, err
}

func (s *Pg) GetEntry(ctx context.Context, id int64) (models.Entry, error) {
	s.logger.Printf("Store. GetEntry was invoked")

	q := `
SELECT id, account_id, amount, created_at FROM entries
	WHERE id = $1 LIMIT 1
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Entry{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	var e models.Entry
	err = row.Scan(&e.ID, &e.AccountID, &e.Amount, &e.CreatedAt)

	// row := s.db.QueryRowContext(ctx, q, id)

	// var e models.Entry
	// err := row.Scan(&e.ID, &e.AccountID, &e.Amount, &e.CreatedAt)

	return e, err
}
