package usecase

import (
	"github.com/MurilojrMarques/api-transaction.git/model"
	"github.com/MurilojrMarques/api-transaction.git/repository"
)

type TransactionUsecase struct {
	repository repository.TransactionRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository) TransactionUsecase {
	return TransactionUsecase{
		repository: repo,
	}
}

func (tu *TransactionUsecase) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	transactionId, err := tu.repository.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction.ID = transactionId

	return transaction, nil
}
