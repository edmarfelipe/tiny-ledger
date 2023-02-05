package create

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
	"github.com/go-chi/chi/v5"
)

func New(transactionDB *db.TransactionDB) utils.APIController {
	return &controller{
		usc: &useCase{
			transactionDB: transactionDB,
		},
	}
}

type controller struct {
	usc *useCase
}

func (ctrl *controller) Handler(w http.ResponseWriter, r *http.Request) (err error) {
	defer utils.Logging("creating transaction", err, time.Now())

	in, err := utils.ParseBody[Input](r.Body)
	if err != nil {
		return err
	}

	in.AccountID = chi.URLParam(r, "account-id")

	out, err := ctrl.usc.Create(r.Context(), *in)
	if err != nil {
		return err
	}

	return utils.SendResponse(w, http.StatusCreated, out)
}
