package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/edmarfelipe/tiny-ledger/entity"
)

func NewPersonDB(conn *sql.DB) *PersonDB {
	return &PersonDB{
		conn: conn,
	}
}

type PersonDB struct {
	conn *sql.DB
}

func (db *PersonDB) Create(ctx context.Context, model entity.Person) error {
	query := `INSERT INTO persons (id, name, doc, birth) VALUES ($1,$2,$3,$4)`
	_, err := db.conn.ExecContext(ctx, query, model.ID, model.Name, model.Doc, model.Birth)
	if err != nil {
		return err
	}

	return nil
}

func (db *PersonDB) Update(ctx context.Context, id string, model entity.Person) error {
	query := `UPDATE persons SET name = $1, doc = $2, birth = $3 WHERE id = $4`
	_, err := db.conn.ExecContext(ctx, query, model.Name, model.Doc, model.Birth, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *PersonDB) FindOne(ctx context.Context, id string) (*entity.Person, error) {
	person := entity.Person{}
	query := "SELECT id, name, birth, doc FROM persons where id = $1"
	err := db.conn.QueryRowContext(ctx, query, id).Scan(&person.ID, &person.Name, &person.Birth, &person.Doc)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &person, nil
}
