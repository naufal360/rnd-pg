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
	env := config.GetEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DB_USER,
		env.DB_PASS,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connected successfully!")

	DB.AutoMigrate(&domain.User{}, &domain.Transaction{})
}
