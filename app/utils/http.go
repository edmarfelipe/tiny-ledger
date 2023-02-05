package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// ParseBody convert json into the specified type
func ParseBody[T any](r io.Reader) (*T, error) {
	body := new(T)
	if err := json.NewDecoder(r).Decode(body); err != nil {
		return nil, err
	}
	return body, nil
}

// URLQueryDate parse date from url, it only accepts the 2006-01-02 format
func URLQueryDate(r *http.Request, key string) (*time.Time, error) {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return nil, nil
	}

	valueDate, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}

	return &valueDate, nil
}

// SendResponse sends http response with an encoded body and status code
func SendResponse(w http.ResponseWriter, statusCode int, body any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(body)
}

// SendError sends http response with an error
func SendError(w http.ResponseWriter, err error) error {
	respError := errors.New("internal server error")
	respStatus := http.StatusInternalServerError

	if _, ok := err.(*DomainError); ok {
		respError = err
		respStatus = http.StatusBadRequest
	}

	return SendResponse(w, respStatus, NewErrorResponse(respError))
}
