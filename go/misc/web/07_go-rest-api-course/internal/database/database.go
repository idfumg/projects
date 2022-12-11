package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	log "github.com/sirupsen/logrus"
)

func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up a new database connection")

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connectStringFmt := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	connectString := fmt.Sprintf(connectStringFmt, dbHost, dbPort, dbUser, dbTable, dbPass, dbSSLMode)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
