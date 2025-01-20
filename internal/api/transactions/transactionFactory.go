package transactions

import (
	"time"
)

func NewTransaction(userID int64, description string, category int, amount float64) *Transaction {
	now := time.Now()
	return &Transaction{
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
		Description: description,
		Category:    category,
		Amount:      amount,
	}
}
