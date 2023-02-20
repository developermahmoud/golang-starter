package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Connection *sql.DB
)

func OpenConnectionToDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("TIME_ZONE"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error To Connect Database")
	}

	return DB
}

func ReturnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func CloseConnection() {
	Connection.Close()
}
