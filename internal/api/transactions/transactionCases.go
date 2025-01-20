package transactions

import (
	"time"
)

type TransactionUseCase interface {
	FilterTransactions(userID int64, filter *Filter) ([]*Transaction, error)
	CreateTransaction(userID int64, description string, category int, amount float64) error
	GetTransactions(userID int64) ([]*Transaction, error)
	UpdateTransaction(transaction *Transaction) error
	DeleteTransaction(userID int64, id int64) error
}

type transactionUseCase struct {
	transactionRepo TransactionRepositories
}

func NewTransactionUseCase(tr TransactionRepositories) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tr,
	}
}

func (uc *transactionUseCase) CreateTransaction(userID int64, description string, category int, amount float64) error {
	transaction := NewTransaction(userID, description, category, amount)
	return uc.transactionRepo.Create(transaction)
}

func (uc *transactionUseCase) GetTransactions(userID int64) ([]*Transaction, error) {
	return uc.transactionRepo.GetAll(userID)
}

func (uc *transactionUseCase) UpdateTransaction(transaction *Transaction) error {
	transaction.UpdatedAt = time.Now()
	return uc.transactionRepo.Update(transaction)
}

func (uc *transactionUseCase) DeleteTransaction(userID int64, id int64) error {
	return uc.transactionRepo.Delete(userID, id)
}

func (uc *transactionUseCase) FilterTransactions(userID int64, filter *Filter) ([]*Transaction, error) {
	return uc.transactionRepo.Filter(userID, filter)
}
