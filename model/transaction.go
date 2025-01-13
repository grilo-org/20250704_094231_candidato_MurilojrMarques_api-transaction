package model

import "time"

type Transaction struct {
	ID          int       `json:"id_transaction"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
}
