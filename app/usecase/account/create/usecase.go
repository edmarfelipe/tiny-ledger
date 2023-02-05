package create

import (
	"context"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/entity"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

type useCase struct {
	accountDB *db.AccountDB
	personDB  *db.PersonDB
}

type Input struct {
	PersonID string `json:"person_id"`
}

func (in *Input) IsValid() bool {
	return len(in.PersonID) > 0
}

type Output struct {
	ID       string    `json:"id"`
	Balance  float64   `json:"balance"`
	PersonID string    `json:"person_id"`
	Date     time.Time `json:"date"`
	Enable   bool      `json:"enable"`
}

func (usc *useCase) Create(ctx context.Context, in Input) (*Output, error) {
	if !in.IsValid() {
		return nil, utils.NewError("input is not valid", utils.InvalidInputError)
	}

	person, err := usc.personDB.FindOne(ctx, in.PersonID)
	if err != nil {
		return nil, err
	}

	if person == nil {
		return nil, utils.NewError("person not found", utils.ResourceFoundError)
	}

	newAccount := entity.NewAccount(in.PersonID)
	err = usc.accountDB.Create(ctx, *newAccount)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:       newAccount.ID,
		Balance:  newAccount.Balance,
		PersonID: newAccount.PersonID,
		Date:     newAccount.Date,
		Enable:   newAccount.Enable,
	}, nil
}
