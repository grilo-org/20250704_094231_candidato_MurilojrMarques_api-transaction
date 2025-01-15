package controller

import (
	"net/http"
	"strconv"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Dados inválidos"})
		return
	}

	insertedTransaction, err := t.transactionUsecase.CreateTransaction(transaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erros": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedTransaction)
}

func (t *transactionController) GetTransactionConverted(ctx *gin.Context) {
	id := ctx.Param("id")
	currency := ctx.DefaultQuery("currency", "USD")

	transactionID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido. Deve ser um número inteiro"})
		return
	}

	transaction, err := t.transactionUsecase.GetTransactionConverted(transactionID, currency)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
