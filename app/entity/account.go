package entity

import (
	"time"

	"github.com/google/uuid"
)

func NewAccount(personID string) *Account {
	return &Account{
		ID:       uuid.New().String(),
		Balance:  0,
		PersonID: personID,
		Date:     time.Now().UTC(),
		Enable:   true,
	}
}

type Account struct {
	ID       string    `json:"id"`
	Balance  float64   `json:"balance"`
	PersonID string    `json:"person_id"`
	Date     time.Time `json:"date"`
	Enable   bool      `json:"enable"`
}
