package controller

import (
	"net/http"

	"github.com/MurilojrMarques/api-transaction.git/model"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(usecase usecase.TransactionUsecase) transactionController {
	return transactionController{
		transactionUsecase: usecase,
	}
}

func (t *transactionController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	err := ctx.BindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTransaction, err := t.transactionUsecase.CreateTransaction(transaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTransaction)
}
