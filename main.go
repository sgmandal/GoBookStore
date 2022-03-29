package main

import (
	"bookstore/api"
	"bookstore/database"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var err error

func main() {
	database.Db, err = sqlx.Connect("mysql", "root:@(localhost:3306)/trial")
	if err != nil {
		log.Println(err)
	}
	defer database.Db.Close()

	// database schema execution, can be put into sifferent function
	database.SeedData()

	r := api.Run()
	r.Run(":8080")
}
