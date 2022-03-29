package database

import (
	"bookstore/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// tried using AUTO_INCREMENT
// wasnt getting desired results
// Query is executed one by one and not completely

// source: https://github.com/jmoiron/sqlx
// exec the schema or fail; multi-statement Exec behavior varies between
// database drivers;  pq will exec them all, sqlite3 won't, ymmv

var schema = []string{
	`
CREATE TABLE Author (
	AuthorID int NOT NULL,
	AuthorName TEXT(255),
	AuthorAge int
);
`,
	`
CREATE TABLE Books (
    BookId int NOT NULL,
	AuthorID int,
    BookName TEXT(100),
    BookUrl TEXT(2048)
);
`,
}

var deleteschema = []string{
	`DROP TABLE Author;`,
	`DROP TABLE Books;`,
}

// var err error

func SeedData() {

	for i, k := range deleteschema {
		_, err := Db.Exec(k)
		if err == nil || err != nil {
			Db.MustExec(schema[i])
		}
		log.Println("changes made")
	}

	Populate()
}

// function to populate the database
func Populate() {
	// temporary data insertion
	Books := []models.Book{
		models.Book{
			Book_id:   1,
			AuthorID:  1,
			Book_Name: "The Alchemist",
			Book_Url:  "https://www.facebook.com",
		},
		models.Book{
			Book_id:   2,
			AuthorID:  2,
			Book_Name: "A Brief History of TIme",
			Book_Url:  "https://en.wikipedia.org/wiki/A_Brief_History_of_Time",
		},
		models.Book{
			Book_id:   3,
			AuthorID:  1,
			Book_Name: "Eleven Minutes",
			Book_Url:  "https://en.wikipedia.org/wiki/Eleven_Minutes",
		},
	}

	Authors := []models.Author{
		models.Author{
			AuthorId:   1,
			AuthorName: "Paulo Coelho",
			AuthorAge:  50,
		},
		models.Author{
			AuthorId:   2,
			AuthorName: "Stephen Hawking",
			AuthorAge:  76,
		},
	}

	for _, book := range Books {
		_, err := Db.NamedExec(`INSERT INTO Books (BookId,AuthorID,BookName,BookUrl) VALUES (:BookId,:AuthorID,:BookName,:BookUrl);`, &book)
		if err != nil {
			log.Println(err)
			return
		}
	}
	for _, author := range Authors {
		_, err := Db.NamedExec(`INSERT INTO Author (AuthorID,AuthorName,AuthorAge) VALUES (:AuthorID,:AuthorName,:AuthorAge);`, &author)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// print and see if the data is added to database correctly
	books1 := []models.Book{}
	authors1 := []models.Author{}
	err := Db.Select(&books1, `SELECT * FROM Books;`)
	if err != nil {
		fmt.Println(err)
	}
	err = Db.Select(&authors1, `SELECT * FROM Author;`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books1)
}
