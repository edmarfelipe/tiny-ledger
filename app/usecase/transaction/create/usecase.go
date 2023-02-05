package create

import (
	"context"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/entity"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

type Input struct {
	Amount    float64 `json:"amount"`
	AccountID string  `json:"-"`
}

func (in *Input) IsValid() bool {
	return len(in.AccountID) > 0
}

type Output struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	AccountID string    `json:"account_id"`
	Date      time.Time `json:"date"`
}

type useCase struct {
	transactionDB *db.TransactionDB
}

func (usc *useCase) Create(ctx context.Context, in Input) (*Output, error) {
	if !in.IsValid() {
		return nil, utils.NewError("input is not valid", utils.InvalidInputError)
	}

	tran := entity.NewTransaction(in.AccountID, in.Amount)
	err := usc.transactionDB.Create(ctx, *tran)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:        tran.ID,
		Amount:    tran.Amount,
		AccountID: tran.AccountID,
		Date:      tran.Date,
	}, nil
}
