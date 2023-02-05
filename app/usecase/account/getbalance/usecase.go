package getbalance

import (
	"context"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

type useCase struct {
	accountDB *db.AccountDB
}

type Input struct {
	AccountID string
}

func (in *Input) IsValid() bool {
	return len(in.AccountID) > 0
}

type Output struct {
	Total float64
}

func (usc *useCase) GetBalance(ctx context.Context, in Input) (*Output, error) {
	if !in.IsValid() {
		return nil, utils.NewError("input is not valid", utils.InvalidInputError)
	}

	account, err := usc.accountDB.FindOne(ctx, in.AccountID)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, utils.NewError("account not found", utils.ResourceFoundError)
	}

	return &Output{Total: account.Balance}, nil
}
