package database

import (
	"bookstore/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

var schema = `
CREATE TABLE Books (
    BookId int,
    BookName TEXT(100),
    BookUrl TEXT(2048)
);
`

var schema2 = `
DROP TABLE Books;
`

// var err error

func SeedData() {
	_, err := Db.Exec(schema2)
	if err == nil || err != nil {
		Db.MustExec(schema)
	}
	log.Println("changes made")

	Populate()
}

// function to populate the database
func Populate() {
	// temporary data insertion
	Books := []models.Book{
		models.Book{
			Book_id:   1,
			Book_Name: "The Alchemist",
			Book_Url:  "https://www.facebook.com",
		},
		models.Book{
			Book_id:   2,
			Book_Name: "A Brief History of TIme",
			Book_Url:  "https://en.wikipedia.org/wiki/A_Brief_History_of_Time",
		},
	}

	for _, book := range Books {
		_, err := Db.NamedExec(`INSERT INTO Books (BookId,BookName, BookUrl) VALUES (:BookId,:BookName, :BookUrl);`, &book)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// print and see if the data is added to database correctly
	books1 := []models.Book{}
	err := Db.Select(&books1, `SELECT * FROM Books`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books1)
}
