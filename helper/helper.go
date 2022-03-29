package helper

import (
	"bookstore/database"
	"bookstore/models"
	"errors"
	"fmt"
)

func GetEveryBook(xd *[]models.Book) error {
	err := database.Db.Select(xd, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books;`) // change
	if err != nil {
		return err
	}
	return nil
}

func GetByAuthor(authorName string, xd *[]models.Book) error {
	// first fetch the author id based on author name
	// based on that author id, return all books
	x := []models.Author{}
	err := database.Db.Select(&x, `SELECT AuthorID FROM author WHERE AuthorName=?;`, authorName)
	if err != nil {
		fmt.Println("fetching author id")
		return err
	}

	err = database.Db.Select(xd, `SELECT * FROM books WHERE AuthorID=?;`, x[0].AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func GetByPublisher(publisherName string, xd *[]models.Book) error {
	// first fetch the author id based on author name
	// based on that author id, return all books
	x := []models.Publisher{}
	err := database.Db.Select(&x, `SELECT PublisherID FROM Publisher WHERE PublisherName=?;`, publisherName)
	if err != nil {
		return err
	}

	fmt.Println(x)
	err = database.Db.Select(xd, `SELECT * FROM books WHERE PublisherID=?;`, x[0].PublisherID)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByIDHelper(id int, book *[]models.Book) error {
	err := database.Db.Select(book, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE BookId=?;`, id)
	if err != nil {
		return err
	} else if len(*book) == 0 {
		return errors.New("Invalid id because no book found")
	}
	return nil
}

/*approach:
dont take id from user
for id, count how many entries are in the database(use count in mysql)

try and set count+1 as new id

use insert into syntax

log errors if any


*/

func PostABookHandler(book *models.Book) error {
	var x []uint8
	err := database.Db.Select(&x, `SELECT COUNT(BookId) AS count FROM books;`)
	if err != nil {
		return err
	}
	fmt.Println("count is:- ", x)
	id := int(x[0]) + 1
	book.Book_id = id

	//named exec

	_, err = database.Db.NamedExec(`INSERT INTO books (BookId,AuthorID,PublisherID,BookName,BookUrl)
	VALUES (:BookId,:AuthorID,:PublisherID,:BookName,:BookUrl);`, book)
	if err != nil {
		return err
	}

	return nil
}

/*approach to update a table

 */
// put update a book helper
func UpdateABookHelper(id int, update *models.Book) error {
	books := []models.Book{}
	err := database.Db.Select(&books, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE BookId=?;`, id)
	if err != nil {
		return err
	} else if len(books) == 0 {
		return errors.New("Invalid id because no book found")
	}

	book := books[0]
	if update.AuthorID != 0 {
		book.AuthorID = update.AuthorID
	} else if update.PublisherID != 0 {
		book.PublisherID = update.PublisherID
	} else if update.Book_Name != "" {
		book.Book_Name = update.Book_Name
	} else if update.Book_Url != "" {
		book.Book_Url = update.Book_Url
	}

	_, err = database.Db.NamedExec(`UPDATE books
	SET BookId=:BookId,AuthorID=:AuthorID, PublisherID=:PublisherID, BookName=:BookName, BookUrl=:BookUrl
	WHERE BookId=:BookId;`, &book)
	if err != nil {
		return err
	}

	// check for missing values in put's struct

	return nil
}
