package model

import "time"

type Transaction struct {
	ID          int64     `json:"id_transaction"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
}
