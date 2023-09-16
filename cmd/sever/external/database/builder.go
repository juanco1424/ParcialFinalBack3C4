package database

import (
	"database/sql"
)

func MySQLDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3307)/dh202302-g4")
	if err != nil {
		panic(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, err
	}

	return db, nil
}
