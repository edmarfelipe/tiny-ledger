package block

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

func (usc *useCase) Block(ctx context.Context, in Input) error {
	if !in.IsValid() {
		return utils.NewError("input is not valid", utils.InvalidInputError)
	}

	account, err := usc.accountDB.FindOne(ctx, in.AccountID)
	if err != nil {
		return err
	}

	if account == nil {
		return utils.NewError("account not found", utils.ResourceFoundError)
	}

	account.Enable = false

	err = usc.accountDB.Update(ctx, in.AccountID, *account)
	if err != nil {
		return err
	}

	return nil
}
