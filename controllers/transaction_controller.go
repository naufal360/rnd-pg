package controllers

import (
	"net/http"
	"payment-gateway/services"
	"payment-gateway/util"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// Endpoint untuk membuat transaksi dengan Midtrans Snap
func (c *TransactionController) CreateTransaction(ctx echo.Context) error {
	type Request struct {
		OrderID       string  `json:"order_id" validate:"required"`
		CustomerName  string  `json:"customer_name" validate:"required"`
		CustomerEmail string  `json:"customer_email" validate:"required,email"`
		GrossAmount   float64 `json:"gross_amount" validate:"required"`
	}

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid input"))
	}

	transaction, err := c.service.CreateTransaction(req.OrderID, req.CustomerName, req.CustomerEmail, req.GrossAmount)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	// Integrasi dengan Midtrans Snap
	// Masukkan kode integrasi di sini
	return ctx.JSON(http.StatusOK, util.NewSuccessResponse(transaction))
}

// Endpoint untuk menerima callback dari Midtrans
func (c *TransactionController) HandleCallback(ctx echo.Context) error {
	type CallbackRequest struct {
		OrderID           string `json:"order_id"`
		TransactionStatus string `json:"transaction_status"`
	}

	var callback CallbackRequest
	if err := ctx.Bind(&callback); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid callback payload"))
	}

	err := c.service.UpdateTransactionStatus(callback.OrderID, callback.TransactionStatus)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, util.NewSuccessResponse("Transaction updated successfully"))
}

func (c *TransactionController) GetTransactions(ctx echo.Context) error {
	transactions, err := c.service.GetTransactions()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}
	return ctx.JSON(http.StatusOK, util.NewSuccessResponse(transactions))
}
