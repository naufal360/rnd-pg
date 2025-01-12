package controllers

import (
	"net/http"
	"payment-gateway/dto"
	"payment-gateway/services"
	"payment-gateway/util"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TransactionController struct {
	logger  *zap.Logger
	service services.TransactionService
}

func NewTransactionController(logger *zap.Logger, service services.TransactionService) *TransactionController {
	return &TransactionController{
		logger:  logger,
		service: service,
	}
}

func (c *TransactionController) CreateTransaction(ctx echo.Context) error {
	var req dto.TransactionRequest
	if err := ctx.Bind(&req); err != nil {
		c.logger.Error("Failed to bind request", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid input"))
	}

	transaction, err := c.service.CreateTransaction(req)
	if err != nil {
		c.logger.Error("Failed to create transaction", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	c.logger.Info("Transaction created successfully", zap.Any("transaction", transaction))

	// Integrasi dengan Midtrans Snap
	// Masukkan kode integrasi di sini
	return ctx.JSON(http.StatusOK, transaction)
}

// Endpoint untuk menerima callback dari Midtrans
func (c *TransactionController) HandleCallback(ctx echo.Context) error {
	type CallbackRequest struct {
		OrderID           string `json:"order_id"`
		TransactionStatus string `json:"transaction_status"`
	}

	var callback CallbackRequest
	if err := ctx.Bind(&callback); err != nil {
		c.logger.Error("Failed to bind callback payload", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid callback payload"))
	}

	err := c.service.UpdateTransactionStatus(callback.OrderID, callback.TransactionStatus)
	if err != nil {
		c.logger.Error("Failed to update transaction status", zap.String("order_id", callback.OrderID), zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	c.logger.Info("Transaction status updated successfully", zap.String("order_id", callback.OrderID), zap.String("status", callback.TransactionStatus))
	return ctx.JSON(http.StatusOK, util.NewSuccessResponse("Transaction updated successfully"))
}

func (c *TransactionController) GetTransactions(ctx echo.Context) error {
	transactions, err := c.service.GetTransactions()
	if err != nil {
		c.logger.Error("Failed to fetch transactions", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	c.logger.Info("Fetched transactions successfully", zap.Int("count", len(transactions)))
	return ctx.JSON(http.StatusOK, util.NewSuccessResponse(transactions))
}
