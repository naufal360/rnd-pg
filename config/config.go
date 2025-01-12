package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

func GetEnv() Env {
	EnvVar := Env{
		DB_USER:            os.Getenv("DB_USER"),
		DB_PASS:            os.Getenv("DB_PASS"),
		DB_HOST:            os.Getenv("DB_HOST"),
		DB_PORT:            os.Getenv("DB_PORT"),
		DB_NAME:            os.Getenv("DB_NAME"),
		PAYMENT_HOST:       os.Getenv("PAYMENT_HOST_URL"),
		PAYMENT_SERVER_KEY: os.Getenv("PAYMENT_SERVER_KEY"),
	}

	return EnvVar
}

type Env struct {
	DB_USER            string
	DB_PASS            string
	DB_HOST            string
	DB_PORT            string
	DB_NAME            string
	PAYMENT_HOST       string
	PAYMENT_SERVER_KEY string
}
