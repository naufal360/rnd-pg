package routes

import (
	"net/http"
	"payment-gateway/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, userController *controllers.UserController, transactionController *controllers.TransactionController) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetUsers) // Get list of users

	e.POST("/transactions", transactionController.CreateTransaction)
	e.POST("/transactions/callback", transactionController.HandleCallback)
	e.GET("/transactions", transactionController.GetTransactions) // Get list of transactions
}
