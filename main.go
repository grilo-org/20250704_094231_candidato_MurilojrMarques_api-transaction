package main

import (
	"log"
	"net/http"

	"github.com/MurilojrMarques/api-transaction.git/controller"
	"github.com/MurilojrMarques/api-transaction.git/database/config"
	"github.com/MurilojrMarques/api-transaction.git/repository"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	postgresDb, err := config.NewPostgresDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer postgresDb.Db.Close()

	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	TransactionRepository := repository.NewTransactionRepository(postgresDb.Db)
	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository)
	TransactionController := controller.NewTransactionController(TransactionUsecase)

	server.POST("/transaction", TransactionController.CreateTransaction)
	server.GET("/transaction/:id/convert", TransactionController.GetTransactionConverted)

	if err := server.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
	}
	log.Println("Servidor rodando na porta 8080")

}
