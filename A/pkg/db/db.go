package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB
}

// New returns a naive implementation of DB.
func New() *DB {
	return &DB{
		DB: Init(),
	}
}

func Init() *sql.DB {
	driverName := os.Getenv("DB_DRIVER_NAME")
	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dataSourceName := user + ":" + password + "@/" + dbName

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err) // If the Open function is failed, then the whole program does not work either.
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
