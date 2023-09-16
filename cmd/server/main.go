package main

import (
	"fmt"
	"parcial/cmd/server/external/database"
)

func main() {

	database.MySQLDatabase()
	fmt.Println("Connection Successful")

}
