package getbalance

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
	"github.com/go-chi/chi/v5"
)

func New(accountDB *db.AccountDB) utils.APIController {
	return &controller{
		usc: &useCase{
			accountDB: accountDB,
		},
	}
}

type controller struct {
	usc *useCase
}

func (ctrl *controller) Handler(w http.ResponseWriter, r *http.Request) error {
	defer utils.Logging("consulting account total", nil, time.Now())

	in := Input{
		AccountID: chi.URLParam(r, "account-id"),
	}

	balance, err := ctrl.usc.GetBalance(r.Context(), in)
	if err != nil {
		return err
	}

	return utils.SendResponse(w, http.StatusOK, balance)
}
