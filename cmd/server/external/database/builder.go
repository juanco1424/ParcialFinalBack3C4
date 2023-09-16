package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3036)/grupo4")
	if err != nil {
		panic(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, err
	}

	return db, nil
}
