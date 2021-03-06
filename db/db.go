package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

//OpenPG ...
func OpenPG() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil

}
