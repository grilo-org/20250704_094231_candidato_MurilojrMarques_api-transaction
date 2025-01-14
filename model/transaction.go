package model

import "time"

type Transaction struct {
	ID          int       `json:"id_transaction"`
	Description string    `json:"description" validate:"required,max=50"`
	Date        time.Time `json:"date" validate:"required"`
	Value       float64   `json:"value" validate:"required,gt=0"`
}
