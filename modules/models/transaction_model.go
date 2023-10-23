package models

import (
	"time"
)

type Transaction struct {
	TransactionID   int       `json:"transaction_id" gorm:"primaryKey;auto_increment:true;index"`
	TransactionDate time.Time `json:"transaction_date" gorm:"type:date"`
	TransactionType string    `json:"transaction_type" gorm:"type:enum('income','expense')"`
	Amount          int       `json:"amount"`
	CategoryID      int       `json:"category_id" gorm:"index"`
	Description     string    `json:"description" gorm:"type:varchar(300)"`
}
