package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/edmarfelipe/tiny-ledger/entity"
)

func NewAccountDB(conn *sql.DB) *AccountDB {
	return &AccountDB{
		conn: conn,
	}
}

type AccountDB struct {
	conn *sql.DB
}

func (db *AccountDB) Create(ctx context.Context, model entity.Account) error {
	query := `INSERT INTO accounts (id, person_id, balance, enable, date) VALUES ($1,$2,$3,$4,$5)`
	_, err := db.conn.ExecContext(ctx, query, model.ID, model.PersonID, model.Balance, model.Enable, model.Date)
	if err != nil {
		return err
	}

	return nil
}

func (db *AccountDB) Update(ctx context.Context, id string, model entity.Account) error {
	query := `UPDATE accounts SET balance = $1, enable = $2, date = $3 WHERE id = $4`
	_, err := db.conn.ExecContext(ctx, query, model.Balance, model.Enable, model.Date, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *AccountDB) FindOne(ctx context.Context, id string) (*entity.Account, error) {
	account := entity.Account{}
	query := "SELECT id, balance, date, enable FROM accounts where id = $1"
	err := db.conn.QueryRowContext(ctx, query, id).Scan(&account.ID, &account.Balance, &account.Date, &account.Enable)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}
