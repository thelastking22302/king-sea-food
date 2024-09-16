package server

import (
	"os"
	"sync"
	"thelastking/kingseafood/pkg/db"
	"thelastking/kingseafood/pkg/logger"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// singleton
type Singleton struct{}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) Run() *gorm.DB {
	myLogger := logger.GetLogger()
	err := godotenv.Load(".env")
	if err != nil {
		myLogger.Errorf("Error loading .env file: %v", err)
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
		myLogger.Errorf("Fails to connect to database: %v", err)
	}
	myLogger.Infof("Connect suscess to database: %v", db)
	return db
}
