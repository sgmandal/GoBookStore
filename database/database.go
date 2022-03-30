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
	PublisherID int,
    BookName TEXT(100),
    BookUrl TEXT(2048),
	IsDeleted int
);
`,
	`CREATE TABLE Publisher (
	PublisherID int,
    PublisherName TEXT(255),
    PublisherAddress TEXT(255)
);
`,
	`
CREATE TABLE Inventory (
	BookId int,
    NoOfBooks int,
    AddedBy TEXT(255),
	IsDeleted int
);
`,
}

var deleteschema = []string{
	`DROP TABLE Author;`,
	`DROP TABLE Books;`,
	`DROP TABLE Publisher;`,
	`DROP TABLE Inventory;`,
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
		{
			Book_id:     1,
			AuthorID:    1,
			PublisherID: 1,
			Book_Name:   "The Alchemist",
			Book_Url:    "https://www.facebook.com",
			IsDeleted:   0,
		},
		{
			Book_id:     2,
			AuthorID:    2,
			PublisherID: 1,
			Book_Name:   "A Brief History of TIme",
			Book_Url:    "https://en.wikipedia.org/wiki/A_Brief_History_of_Time",
			IsDeleted:   0,
		},
		{
			Book_id:     3,
			AuthorID:    1,
			PublisherID: 1,
			Book_Name:   "Eleven Minutes",
			Book_Url:    "https://en.wikipedia.org/wiki/Eleven_Minutes",
			IsDeleted:   0,
		},
	}

	Authors := []models.Author{
		{
			AuthorId:   1,
			AuthorName: "Paulo Coelho",
			AuthorAge:  50,
		},
		{
			AuthorId:   2,
			AuthorName: "Stephen Hawking",
			AuthorAge:  76,
		},
	}

	PublisherAddress := []models.Publisher{
		{
			PublisherID:      1,
			PublisherName:    "Oxford",
			PublisherAddress: "Oxford, United Kingdom",
		},
	}

	Inventorys := []models.Inventory{
		{
			BookId:    2,
			NoOfBooks: 10,
			AddedBy:   "Sudhakar",
			IsDeleted: 0,
		},
	}
	for _, book := range Books {
		_, err := Db.NamedExec(`INSERT INTO Books (BookId,AuthorID,PublisherID,BookName,BookUrl,IsDeleted) VALUES (:BookId,:AuthorID,:PublisherID,:BookName,:BookUrl,:IsDeleted);`, &book)
		if err != nil {
			fmt.Println("updating books")
			log.Println(err)
			return
		}
	}
	for _, author := range Authors {
		_, err := Db.NamedExec(`INSERT INTO Author (AuthorID,AuthorName,AuthorAge) VALUES (:AuthorID,:AuthorName,:AuthorAge);`, &author)
		if err != nil {
			fmt.Println("updating authors")
			log.Println(err)
			return
		}
	}
	for _, publisher := range PublisherAddress {
		_, err := Db.NamedExec(`INSERT INTO Publisher (PublisherID,PublisherName,PublisherAddress) VALUES (:PublisherID,:PublisherName,:PublisherAddress);`, &publisher)
		if err != nil {
			fmt.Println("updating publisher")
			log.Println(err)
			return
		}
	}
	for _, publisher := range Inventorys {
		_, err := Db.NamedExec(`INSERT INTO Inventory (BookId,NoOfBooks,AddedBy,IsDeleted) VALUES (:BookId,:NoOfBooks,:AddedBy,:IsDeleted);`, &publisher)
		if err != nil {
			fmt.Println("updating Inventory")
			log.Println(err)
			return
		}
	}

	// print and see if the data is added to database correctly
	books1 := []models.Book{}
	authors1 := []models.Author{}
	publisher1 := []models.Publisher{}
	err := Db.Select(&books1, `SELECT * FROM Books;`)
	if err != nil {
		fmt.Println(err)
	}
	err = Db.Select(&authors1, `SELECT * FROM Author;`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books1)
	err = Db.Select(&publisher1, `SELECT * FROM Publisher;`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(publisher1)

	var x []uint8
	err = Db.Select(&x, `SELECT COUNT(BookId) AS count FROM books;`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("count:- ", x)

	// testing put statement
	// test := models.Book{
	// 	Book_Name: "Theory of everything",
	// }
	// y := 2
	// books := []models.Book{}
	// err = Db.Select(&books, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE BookId=?;`, y)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else if len(books) == 0 {
	// 	fmt.Println("no book found")
	// 	return
	// }
	// book := books[0]
	// fmt.Println(book)
	// fmt.Println(test)
	// if test.AuthorID != 0 {
	// 	book.AuthorID = test.AuthorID
	// } else if test.PublisherID != 0 {
	// 	book.PublisherID = test.PublisherID
	// } else if test.Book_Name != "" {
	// 	book.Book_Name = test.Book_Name
	// } else if test.Book_Url != "" {
	// 	book.Book_Url = test.Book_Url
	// }

	// fmt.Println(test)
	// fmt.Println(book)

	// // is transaction required here
	// _, err = Db.NamedExec(`UPDATE books
	// SET BookId=:BookId,AuthorID=:AuthorID, PublisherID=:PublisherID, BookName=:BookName, BookUrl=:BookUrl
	// WHERE BookId=:BookId;`, &book)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	var bookid int = 2
	var NoOfBooks []uint8
	err = Db.Select(&NoOfBooks, `SELECT NoOfBooks FROM inventory
	WHERE BookId=?;`, bookid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(NoOfBooks)
}
