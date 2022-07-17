package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() error {
	err := godotenv.Load("../config/.env")
	if err != nil {
		fmt.Printf("Error loading .nv file: %s\n", err.Error())
		return err
	}
	return nil
}
