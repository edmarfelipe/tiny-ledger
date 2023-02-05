package create

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/utils"
)

func New(personDB *db.PersonDB) utils.APIController {
	return &controller{
		usc: &useCase{
			personDB: personDB,
		},
	}
}

type controller struct {
	usc *useCase
}

func (ctrl *controller) Handler(w http.ResponseWriter, r *http.Request) (err error) {
	defer utils.Logging("creating person", err, time.Now())

	in, err := utils.ParseBody[Input](r.Body)
	if err != nil {
		return err
	}

	out, err := ctrl.usc.Create(r.Context(), *in)
	if err != nil {
		return err
	}

	return utils.SendResponse(w, http.StatusCreated, out)
}
