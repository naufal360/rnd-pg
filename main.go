package main

import (
	"log"
	"payment-gateway/config"
	"payment-gateway/connection"
	"payment-gateway/controllers"
	"payment-gateway/provider"
	"payment-gateway/repo"
	"payment-gateway/routes"
	"payment-gateway/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	config.LoadEnv()
	connection.ConnectDB()

	// provider
	midtransProvider := provider.NewMidtrans()

	userRepo := repo.NewUserRepository(connection.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	transactionRepo := repo.NewTransactionRepository(connection.DB)
	transactionService := services.NewTransactionService(transactionRepo, midtransProvider)
	transactionController := controllers.NewTransactionController(logger, transactionService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	// Add Logger Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e, userController, transactionController)

	e.Logger.Fatal(e.Start(":8000"))
}
