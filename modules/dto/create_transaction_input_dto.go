package dto

import "time"

type CreateTransactionInput struct {
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	TransactionType string    `json:"transaction_type" validate:"required"`
	Amount          int       `json:"amount" validate:"required"`
	CategoryID      int       `json:"category_id" validate:"required"`
	Description     string    `json:"description" validate:"required"`
}
