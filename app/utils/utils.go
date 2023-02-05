package utils

import (
	"log"
	"net/http"
	"time"
)

func Logging(msg string, err error, start time.Time) {
	log.Printf("msg=%s err=%v took=%s \n", msg, err, time.Since(start))
}

type APIController interface {
	Handler(w http.ResponseWriter, r *http.Request) (err error)
}
