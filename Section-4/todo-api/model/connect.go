package model

import (
	"database/sql"
	"fmt"
	"log"
)

//use only simple c not C because if C add then all packages can access the DB, but use c it only model package can do that

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@/tcp:localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db

}
