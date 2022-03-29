package helper

import (
	"bookstore/database"
	"bookstore/models"
	"fmt"
)

func GetEveryBook(xd *[]models.Book) error {
	err := database.Db.Select(xd, `SELECT * FROM Books;`) // change
	if err != nil {
		return err
	}
	return nil
}

func GetByAuthor(authorName string, xd *[]models.Book) error {
	// first fetch the author id based on author name
	// based on that author id, return all books
	x := []models.Author{}
	err := database.Db.Select(&x, `SELECT AuthorID FROM author WHERE AuthorName=?`, authorName)
	if err != nil {
		fmt.Println("fetching author id")
		return err
	}

	err = database.Db.Select(xd, `SELECT * FROM books WHERE AuthorID=?`, x[0].AuthorId)
	if err != nil {
		return err
	}
	return nil
}
