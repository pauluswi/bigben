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

func (repository *ewalletRepositoryImpl) Transfer(fromAccountNumber int32, toAccountNumber int32, amount int) (err error) {
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

	query := `UPDATE db_wallet.wallet SET balance = ?, modifieddate = Now() WHERE accountid = ?`

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

	insert := `INSERT INTO db_wallet.wallet_trx(accountid, trxtype, dc, trxamount, createdby, createddate, modifiedby, modifieddate) VALUES (?, ?, ?, ?, ?, Now(), ?, Now())`
	_, err = tx.ExecContext(ctx, insert, fromAccountNumber, "P2P TRANSFER", "D", amount, "user", "user")
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, insert, fromAccountNumber, "P2P TRANSFER", "C", amount, "user", "user")
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

func (repository *ewalletRepositoryImpl) Find(accountNumber int32) (balance *entity.Ewallet, err error) {
	var (
		accountNumberRow sql.NullInt32
		balanceRow       sql.NullInt32
		modifiedDateRow  sql.NullString
	)
	ctx, cancel := config.NewMySQLContext()
	defer cancel()

	query := `SELECT accountid, balance, modifieddate FROM db_wallet.wallet WHERE accountid = ?`
	err = repository.Database.QueryRowContext(ctx, query, accountNumber).Scan(
		&accountNumberRow,
		&balanceRow,
		&modifiedDateRow,
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
	balance = &entity.Ewallet{
		AccountID:    accountNumberRow.Int32,
		Balance:      int(balanceRow.Int32),
		ModifiedDate: modifiedDateRow.String,
	}
	return balance, nil
}

func (repository *ewalletRepositoryImpl) FindTransactions(accountNumber int32) (trxList []entity.EwalletTrx, err error) {
	ctx, cancel := config.NewMySQLContext()
	defer cancel()

	query := `SELECT trxid, accountid, trxtype, dc, trxamount, createdby, createddate, modifiedby, modifieddate FROM db_wallet.wallet_trx WHERE accountid = ?`

	rows, err := repository.Database.QueryContext(ctx, query, accountNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			// throw err account number not found
			err = errors.New("account number not found")
			return nil, err
		}
		return nil, err
	}
	defer rows.Close()

	var trxLists []entity.EwalletTrx
	//var trxLists1 []entity.EwalletTrx

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var trxList entity.EwalletTrx
		if err := rows.Scan(&trxList.TrxID, &trxList.AccountID, &trxList.TrxType,
			&trxList.CD, &trxList.Amount, &trxList.CreatedBy, &trxList.CreatedDate,
			&trxList.ModifiedBy, &trxList.ModifiedDate); err != nil {
			return trxLists, err
		}
		trxLists = append(trxLists, trxList)
	}
	return trxLists, nil
}
