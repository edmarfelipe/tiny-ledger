package gettransactions

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
	"github.com/go-chi/chi/v5"
)

func New(tranDB *db.TransactionDB) utils.APIController {
	return &controller{
		usc: &useCase{
			transactionDB: tranDB,
		},
	}
}

type controller struct {
	usc *useCase
}

func (ctrl *controller) Handler(w http.ResponseWriter, r *http.Request) (err error) {
	defer utils.Logging("consulting statement", err, time.Now())

	in := Input{
		AccountID: chi.URLParam(r, "account-id"),
	}

	in.Begin, err = utils.URLQueryDate(r, "begin")
	if err != nil {
		return utils.NewError("url parameters is invalid", utils.InvalidInputError)
	}

	in.End, err = utils.URLQueryDate(r, "end")
	if err != nil {
		return utils.NewError("url parameters is invalid", utils.InvalidInputError)
	}

	transactions, err := ctrl.usc.Find(r.Context(), in)
	if err != nil {
		return err
	}

	return utils.SendResponse(w, http.StatusOK, transactions)
}
