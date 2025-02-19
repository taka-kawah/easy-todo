package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbDsn() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
}

func GetDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDbDsn()))
	if err != nil {
		log.Fatal("Error Connecting to PostgreSQL")
	} else {
		log.Print("Success Connecting to PostgreSQL!")
	}
	return db
}
