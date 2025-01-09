package connection

import (
	"fmt"
	"log"
	"payment-gateway/config"
	"payment-gateway/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASS"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connected successfully!")

	DB.AutoMigrate(&domain.User{}, &domain.Transaction{})
}
