package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ConfigInit() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}
