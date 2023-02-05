package gettransactions

import (
	"context"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/entity"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

type useCase struct {
	transactionDB *db.TransactionDB
}

type Input struct {
	Begin     *time.Time `json:"-"`
	End       *time.Time `json:"-"`
	AccountID string     `json:"-"`
}

func (in *Input) IsValid() bool {
	return len(in.AccountID) > 0
}

func (usc *useCase) Find(ctx context.Context, in Input) ([]entity.Transaction, error) {
	if !in.IsValid() {
		return nil, utils.NewError("input is not valid", utils.InvalidInputError)
	}

	transactions, err := usc.transactionDB.FindAll(ctx, in.AccountID, in.Begin, in.End)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
