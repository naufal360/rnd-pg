package dto

type TransactionRequest struct {
	CustomerFirstName string  `json:"customer_firstname" validate:"required"`
	CustomerLastName  string  `json:"customer_lastname" validate:"required"`
	CustomerEmail     string  `json:"customer_email" validate:"required,email"`
	GrossAmount       float64 `json:"gross_amount" validate:"required"`
}
