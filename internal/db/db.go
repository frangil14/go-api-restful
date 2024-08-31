package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(user, password, dbname, dbport string) error {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, dbport, dbname)
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    if err := DB.Ping(); err != nil {
        return fmt.Errorf("error connecting to the database: %w", err)
    }

    fmt.Println("Successfully connected to the database!")
    return nil
}