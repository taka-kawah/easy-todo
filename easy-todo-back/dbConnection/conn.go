package dbConnection

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbDsn() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: ", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Tokyo sslmode=disable", host, user, password, dbname, port), nil
}

func ConnectToDb() (*gorm.DB, *sql.DB, error) {
	dbDsn, err := getDbDsn()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, nil, fmt.Errorf("error Connecting to PostgreSQL: %w", err)
	}

	sqldb, err := getSQLDB(db)
	if err != nil {
		return db, nil, err
	}

	return db, sqldb, nil
}

func getSQLDB(db *gorm.DB) (*sql.DB, error) {
	sqldb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error Getting SQLInstance: %w", err)
	}
	return sqldb, nil
}

func DisconnectToDb(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error Disconnecting to PostgreSQL: %w", err)
	}
	return nil
}
