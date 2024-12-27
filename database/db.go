package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "postgres://postgres:12345678@localhost/go_crud?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println(("Suceessfully conncted to the Database"))
	return db, nil

}
