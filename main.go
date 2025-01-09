package main

import (
	"payment-gateway/config"
	"payment-gateway/connection"
	"payment-gateway/controllers"
	"payment-gateway/repo"
	"payment-gateway/routes"
	"payment-gateway/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()
	connection.ConnectDB()

	userRepo := repo.NewUserRepository(connection.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	transactionRepo := repo.NewTransactionRepository(connection.DB)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	e := echo.New()

	// Add Logger Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e, userController, transactionController)

	e.Logger.Fatal(e.Start(":8080"))
}
