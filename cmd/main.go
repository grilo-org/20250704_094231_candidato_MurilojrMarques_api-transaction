package main

import (
	"github.com/MurilojrMarques/api-transaction.git/database/config"
	"github.com/MurilojrMarques/api-transaction.git/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	postgresDb, err := config.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	TransactionRepository := repository.NewTransactionRepository(postgresDb.Db)

	server.POST("/transaction")

	server.Run(":8080")
}
