package store

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/models"

	"github.com/jmoiron/sqlx"
)

func (s *Pg) execTx(ctx context.Context, fn func(s *Pg) error) error {
	tx, err := s.db.(*sqlx.DB).BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}

	stx := &Pg{
		db:     tx,
		logger: s.logger,
		config: s.config,
	}

	err = fn(stx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

func (s *Pg) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	s.logger.Printf("Store. CreateEntry was invoked")

	var result TransferTxResult

	err := s.execTx(ctx, func(s *Pg) error {
		// err := func() error {
		var err error
		result.Transfer, err = s.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			s.logger.Printf("CreateTransfer: %v", err)
			return err
		}

		result.FromEntry, err = s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			s.logger.Printf("CreateEntry: %v", err)
			return err
		}

		result.ToEntry, err = s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			s.logger.Printf("CreateEntry: %v", err)
			return err
		}

		result.FromAccount, err = s.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			s.logger.Printf("GetAccountForUpdate: %v", err)
			return err
		}
		result.FromAccount.Balance -= arg.Amount

		result.FromAccount, err = s.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: result.FromAccount.Balance,
		})
		if err != nil {
			s.logger.Printf("UpdateAccount: %v", err)
			return err
		}

		result.ToAccount, err = s.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			s.logger.Printf("GetAccountForUpdate: %v", err)
			return err
		}
		result.ToAccount.Balance += arg.Amount

		result.ToAccount, err = s.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: result.ToAccount.Balance,
		})
		if err != nil {
			s.logger.Printf("UpdateAccount: %v", err)
			return err
		}

		return nil
		// }()
	})

	return result, err
}
