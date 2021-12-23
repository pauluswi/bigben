package repository

import (
	"database/sql"
	"errors"

	"github.com/pauluswi/bigben/config"
	"github.com/pauluswi/bigben/entity"
	"github.com/pauluswi/bigben/exception"
)

func NewEWalletRepository(database *sql.DB) EWalletRepository {
	return &ewalletRepositoryImpl{
		Database: database,
	}
}

type ewalletRepositoryImpl struct {
	Database *sql.DB
}

func (repository *ewalletRepositoryImpl) Update(fromAccountNumber int32, toAccountNumber int32, amount int) (err error) {
	ctx, cancel := config.NewMySQLContext()
	defer cancel()

	tx, err := repository.Database.BeginTx(ctx, nil)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	// Check balance sender
	// if balance < amount = throw err
	// do update source balance -=
	// and do update destinaton balance +=
	// any err do rollback

	currentBalance, err := repository.Find(fromAccountNumber)
	if err != nil {
		return err
	}

	if currentBalance.Balance < amount {
		err := errors.New("insufficient balance")
		return exception.ValidationError{Message: err.Error()}
	}

	deductBalance := (currentBalance.Balance - amount)

	query := `UPDATE account SET balance = ? WHERE account_number = ?`

	_, err = tx.ExecContext(ctx, query, deductBalance, fromAccountNumber)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	receiverBalance, err := repository.Find(toAccountNumber)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	additionBalance := (receiverBalance.Balance + amount)
	_, err = tx.ExecContext(ctx, query, additionBalance, toAccountNumber)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		if errRolled := tx.Rollback(); errRolled != nil {
			return errRolled
		}
		return err
	}

	return nil

}

func (repository *ewalletRepositoryImpl) Find(accountNumber int32) (balance *entity.Balance, err error) {
	var (
		accountNumberRow sql.NullInt32
		accountNameRow   sql.NullString
		balanceRow       sql.NullInt32
	)
	ctx, cancel := config.NewMySQLContext()
	defer cancel()

	query := `SELECT a.account_number, c.name, a.balance FROM customer c INNER JOIN account a ON c.customer_number = a.customer_number WHERE a.account_number = ?`
	err = repository.Database.QueryRowContext(ctx, query, accountNumber).Scan(
		&accountNumberRow,
		&accountNameRow,
		&balanceRow,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// throw err account number not found
			err = errors.New("account number not found")
			return nil, err
		}
		return nil, err
	}

	// no err found assign to struct
	balance = &entity.Balance{
		AccountNumber: accountNumberRow.Int32,
		Name:          accountNameRow.String,
		Balance:       int(balanceRow.Int32),
	}
	return balance, nil
}
