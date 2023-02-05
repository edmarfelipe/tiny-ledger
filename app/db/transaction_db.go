package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/edmarfelipe/tiny-ledger/entity"
)

func NewTransactionDB(conn *sql.DB) *TransactionDB {
	return &TransactionDB{
		conn: conn,
	}
}

type TransactionDB struct {
	conn *sql.DB
}

func (db *TransactionDB) Create(ctx context.Context, model entity.Transaction) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO transactions (id, amount, account_id, date) VALUES ($1, $2, $3, $4)`
	_, err = db.conn.ExecContext(ctx, query, model.ID, model.Amount, model.AccountID, model.Date)
	if err != nil {
		return err
	}

	queryAccount := "UPDATE accounts SET balance = balance + $1 WHERE id = $2"
	_, err = db.conn.ExecContext(ctx, queryAccount, model.Amount, model.AccountID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *TransactionDB) FindAll(ctx context.Context, accountID string, begin *time.Time, end *time.Time) ([]entity.Transaction, error) {
	var trans []entity.Transaction

	if begin == nil {
		date := time.Now().AddDate(0, -1, 0)
		begin = &date
	}

	if end == nil {
		date := time.Now()
		end = &date
	}

	query := `
		select id, amount, date, account_id 
		from transactions 
		where account_id = $1
		and date::date >= $2
		and date::date <= $3
	`

	rows, err := db.conn.QueryContext(ctx, query, accountID, begin.Format("2006-01-02"), end.Format("2006-01-02"))
	if err != nil {
		return trans, err
	}
	defer rows.Close()

	for rows.Next() {
		var tran entity.Transaction
		err := rows.Scan(&tran.ID, &tran.Amount, &tran.Date, &tran.AccountID)
		if err != nil {
			return []entity.Transaction{}, nil
		}
		trans = append(trans, tran)
	}

	return trans, nil
}
