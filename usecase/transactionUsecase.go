package usecase

import (
	"fmt"
	"math"

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
	if len(transaction.Description) > 50 {
		return model.Transaction{}, fmt.Errorf("a descrição não pode exceder 50 caracteres")
	}

	if transaction.Value <= 0 {
		return model.Transaction{}, fmt.Errorf("o valor da compra deve ser positivo")
	}

	transaction.Value = math.Round(transaction.Value*100) / 100

	transactionId, err := tu.repository.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("erro ao criar transação: %s", err.Error())
	}

	transaction.ID = transactionId

	return transaction, nil
}
