package domain

import "time"

type Transaction struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	OrderID           string    `json:"order_id"`
	GrossAmount       float64   `json:"gross_amount"`
	PaymentType       string    `json:"payment_type"`
	TransactionTime   time.Time `json:"transaction_time"`
	TransactionStatus string    `json:"transaction_status"`
	CustomerName      string    `json:"customer_name"`
	CustomerEmail     string    `json:"customer_email"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
