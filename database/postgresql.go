package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() (*sql.DB, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Printf("Error loading .nv file: %s\n", err.Error())
		return nil, err
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return nil, err
	}
	fmt.Println("DB is open")

	if err := db.Ping(); err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return nil, err
	}
	fmt.Println("DB pinged succesfully")

	return db, nil
}

func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
