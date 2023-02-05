package entity

import (
	"time"

	"github.com/google/uuid"
)

func NewTransaction(accountID string, amount float64) *Transaction {
	return &Transaction{
		ID:        uuid.New().String(),
		Amount:    amount,
		AccountID: accountID,
		Date:      time.Now().UTC(),
	}
}

type Transaction struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	AccountID string    `json:"account_id"`
	Date      time.Time `json:"date"`
}
