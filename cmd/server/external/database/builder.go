package database

import (
	"database/sql"
	"fmt"
	"log"
)

func NewMySQLDatabase() (*sql.DB, error) {
	// INIT DB
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dh202302g4")
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3307)/dh202302-g2")
	if err != nil {
		log.Fatal(err)
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing)
	}

	//defer db.Close()

	// Crear tabla paciente
	createTableSQL := `
	    CREATE TABLE IF NOT EXISTS patient (
	        ID INT AUTO_INCREMENT PRIMARY KEY,
	        Name VARCHAR(255) NOT NULL,
	        LastName VARCHAR(255) NOT NULL,
	        Address VARCHAR(255) NOT NULL,
	        DNI VARCHAR(20) NOT NULL,
	        DischargeDate DATE
	    )`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabla paciente creada.")

	insertSQL := "INSERT INTO patient (Name, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"

	// Prepare the SQL statement.
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Close the prepared statement when done.

	result, err := stmt.Exec("John", "Doe", "123 Main St", "1234567890", "2023-01-15")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %d rows.\n", rowsAffected)

	// Read the SQL script from the file
	/*sqlScript, err := os.ReadFile("../init.sql")
		if err != nil {
			log.Fatal(err)
		}

	// Execute the SQL script
	/*_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println("Database initialized successfully")
	return db, nil
}
