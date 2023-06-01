// Package banshee is the starting point for the Lugh project
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"

type config struct {
	// add an http server
	// add a db client
}

func main() {
	fmt.Println("hello")
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("goodbye")
}
