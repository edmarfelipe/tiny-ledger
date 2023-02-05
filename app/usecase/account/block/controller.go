package block

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	usc *useCase
}

func New(accountDB *db.AccountDB) utils.APIController {
	return &controller{
		usc: &useCase{
			accountDB: accountDB,
		},
	}
}

func (ctrl *controller) Handler(w http.ResponseWriter, r *http.Request) (err error) {
	defer utils.Logging("blocking account", err, time.Now())

	in := Input{
		AccountID: chi.URLParam(r, "account-id"),
	}

	err = ctrl.usc.Block(r.Context(), in)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write([]byte("Account updated"))
	return nil
}
