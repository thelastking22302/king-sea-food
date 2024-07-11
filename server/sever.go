package server

import (
	"fmt"
	"log"
	"os"
	"thelastking/kingseafood/pkg/db"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Run() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userName := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	config := &db.Config{
		Host:     host,
		Port:     port,
		Password: password,
		User:     userName,
		DbName:   dbName,
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatalf("Fails to connect to database: %v", err)
	}
	fmt.Printf("Connect suscess to database: %v", db)
	return db
}
