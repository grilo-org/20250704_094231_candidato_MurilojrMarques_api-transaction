package main

import (
	"github.com/MurilojrMarques/api-transaction.git/controller"
	"github.com/MurilojrMarques/api-transaction.git/database/config"
	"github.com/MurilojrMarques/api-transaction.git/repository"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	postgresDb, err := config.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	TransactionRepository := repository.NewTransactionRepository(postgresDb.Db)
	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository)
	TransactionController := controller.NewTransactionController(TransactionUsecase)

	server.POST("/transaction", TransactionController.CreateTransaction)

	server.Run(":8080")
}
