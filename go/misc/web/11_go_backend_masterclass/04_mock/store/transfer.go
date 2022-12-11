package store

import (
	"context"
	"myapp/models"
)

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (s *Pg) CreateTransfer(ctx context.Context, arg CreateTransferParams) (models.Transfer, error) {
	s.logger.Printf("Store. CreateTransfer was invoked")

	q := `
INSERT INTO transfers (
	from_account_id,
	to_account_id,
	amount
) VALUES (
	$1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Transfer{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var t models.Transfer
	err = row.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount, &t.CreatedAt)

	// row := s.db.QueryRowContext(ctx, q, arg.FromAccountID, arg.ToAccountID, arg.Amount)

	// var t models.Transfer
	// err := row.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount, &t.CreatedAt)

	return t, err
}

func (s *Pg) GetTransfer(ctx context.Context, id int64) (models.Transfer, error) {
	s.logger.Printf("Store. GetTransfer was invoked")

	q := `
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
	WHERE id = $1 LIMIT 1
`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return models.Transfer{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	var t models.Transfer
	err = row.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount, &t.CreatedAt)

	// row := s.db.QueryRowContext(ctx, q, id)

	// var t models.Transfer
	// err := row.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount, &t.CreatedAt)

	return t, err
}
