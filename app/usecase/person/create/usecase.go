package create

import (
	"context"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/entity"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

type useCase struct {
	personDB *db.PersonDB
}

type Input struct {
	Name  string    `json:"name"`
	Doc   string    `json:"doc"`
	Birth time.Time `json:"birth"`
}

type Output struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Doc   string    `json:"doc"`
	Birth time.Time `json:"birth"`
}

func (in *Input) IsValid() bool {
	return len(in.Name) > 0 && len(in.Doc) > 0 && !in.Birth.IsZero()
}

func (usc *useCase) Create(ctx context.Context, in Input) (*Output, error) {
	if !in.IsValid() {
		return nil, utils.NewError("input is not valid", utils.InvalidInputError)
	}

	newPerson := entity.NewPerson(in.Name, in.Doc, in.Birth)
	err := usc.personDB.Create(ctx, *newPerson)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:    newPerson.ID,
		Name:  newPerson.Name,
		Doc:   newPerson.Doc,
		Birth: newPerson.Birth,
	}, nil
}
