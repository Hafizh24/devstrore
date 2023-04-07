package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)



func ConnectDB(DBDriver string, DBConn string) (*sqlx.DB, error) {

	db,err := sqlx.Connect(DBDriver, DBConn)
	if err != nil {
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		return nil,err
	}

	log.Print("Database connection established.")

	return db, nil
}