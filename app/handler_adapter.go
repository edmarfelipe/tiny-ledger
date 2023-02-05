package main

import (
	"net/http"

	"github.com/edmarfelipe/tiny-ledger/utils"
)

func handle(ctrl utils.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ctrl.Handler(w, r)
		if err != nil {
			_ = utils.SendError(w, err)
			return
		}
	}
}
