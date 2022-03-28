package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE Person (
    FirstName TEXT(100),
    LastName TEXT(100)
);
`

var schema2 = `
DROP TABLE Person;
`

type Person struct {
	First_Name string `db:"FirstName"`
	Last_Name  string `db:"LastName"`
}

func main() {
	// dbUser:dbPassword@(dbURL:PORT/dbName)
	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/trial")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database connected successfully")
	}
	defer db.Close()

	// create a new instance of database everytime we run our api
	_, err = db.Exec(schema2)
	if err == nil || err != nil {
		db.MustExec(schema)
	}
	fmt.Println("changes made")

	// custom changes
	SeedData := Person{
		First_Name: "Golang",
		Last_Name:  "Playground",
	}

	db.MustExec(`INSERT INTO Person (FirstName, LastName) VALUES (?, ?);`, "Jason", "Moiron")
	db.NamedExec(`INSERT INTO Person (FirstName, LastName) VALUES (:FirstName, :LastName);`, &SeedData)

	persons := []Person{}
	err = db.Select(&persons, "SELECT * FROM Person;")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(persons)

}
