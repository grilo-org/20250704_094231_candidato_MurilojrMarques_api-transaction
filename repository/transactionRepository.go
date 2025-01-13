package repository

import (
	"database/sql"
	"fmt"

	"github.com/MurilojrMarques/api-transaction.git/model"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (tr *TransactionRepository) CreateTransaction(transaction model.Transaction) (int, error) {

	var id int
	query, err := tr.connection.Prepare("INSERT INTO transaction" +
		"(description, date, value)" +
		"VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(transaction.Description, transaction.Date, transaction.Value).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
