package helper

import (
	"bookstore/database"
	"bookstore/models"
	"errors"
	"fmt"
)

// delete changes added
func GetEveryBook(xd *[]models.Book) error {
	err := database.Db.Select(xd, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE IsDeleted=0;`) // change
	if err != nil {
		return err
	}
	return nil
}

// delete changes added
func GetByAuthor(authorName string, xd *[]models.Book) error {
	// first fetch the author id based on author name
	// based on that author id, return all books
	x := []models.Author{}
	err := database.Db.Select(&x, `SELECT AuthorID FROM author WHERE AuthorName=?;`, authorName)
	if err != nil {
		fmt.Println("fetching author id")
		return err
	}

	err = database.Db.Select(xd, `SELECT * FROM books WHERE AuthorID=? AND IsDeleted=0;`, x[0].AuthorId)
	if err != nil {
		return err
	}
	return nil
}

// delete changes added
func GetByPublisher(publisherName string, xd *[]models.Book) error {
	// first fetch the author id based on author name
	// based on that author id, return all books
	x := []models.Publisher{}
	err := database.Db.Select(&x, `SELECT PublisherID FROM Publisher WHERE PublisherName=?;`, publisherName)
	if err != nil {
		return err
	}

	fmt.Println(x)
	err = database.Db.Select(xd, `SELECT * FROM books WHERE PublisherID=? AND IsDeleted=0;`, x[0].PublisherID)
	if err != nil {
		return err
	}
	return nil
}

// delete changes added
func GetBookByIDHelper(id int, book *[]models.Book) error {

	err := database.Db.Select(book, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE BookId=? AND IsDeleted=0;`, id)
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

// delete changes added
func PostABookHandler(book *models.Book) error {
	var x []uint8
	err := database.Db.Select(&x, `SELECT COUNT(BookId) AS count FROM books WHERE IsDeleted=0;`)
	if err != nil {
		return err
	}
	fmt.Println("count is:- ", x)
	id := int(x[0]) + 1
	book.Book_id = id
	book.IsDeleted = 0

	//named exec

	_, err = database.Db.NamedExec(`INSERT INTO books (BookId,AuthorID,PublisherID,BookName,BookUrl,IsDeleted)
	VALUES (:BookId,:AuthorID,:PublisherID,:BookName,:BookUrl,:IsDeleted);`, book)
	if err != nil {
		return err
	}

	return nil
}

// delete changes added
func UpdateABookHelper(id int, update *models.Book) error {
	books := []models.Book{}
	err := database.Db.Select(&books, `SELECT BookId, AuthorID, PublisherID, BookName, BookUrl FROM Books WHERE BookId=? AND IsDeleted=0;`, id)
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

// delete changes added
func ShowInventoryHelper(inventory *[]models.Inventory) error {
	err := database.Db.Select(inventory, `Select BookId, NoOfBooks, AddedBy FROM Inventory WHERE IsDeleted=0;`)
	if err != nil {
		return err
	}
	return nil
}

// delete changes added
func PurchaseABookHelper(bookid int) error {
	/*approach:
	get book id
	fetch inventory i.e., no of books with bookid
	decrement the inventory [uint8]
	update databse using this decremented inventory number[use transaction concept here]
	optional: return updated inventory in handler
	*/
	var NoOfBooksList []uint8
	err := database.Db.Select(&NoOfBooksList, `SELECT NoOfBooks FROM inventory
	WHERE BookId=? AND IsDeleted=0;`, bookid)
	if err != nil {
		return err
	}
	NoOfBooks := int(NoOfBooksList[0] - 1)

	// start transaction
	tx, err := database.Db.Begin()
	if err != nil {
		return err
	}
	tx.Exec(`UPDATE inventory
	SET NoOfBooks=?
	WHERE BookId=?;`, NoOfBooks, bookid)
	tx.Commit()

	return nil
}

func DeleteABookHelper(bookid int) error {
	// soft delete records from both inventory and books
	// add where clause in purchase a book(done), show inventory(done), update a book, get book by id, get by publisher, get by author, get every book

	// updating inventory
	_, err := database.Db.Exec(`UPDATE inventory
	SET IsDeleted=1
	WHERE BookId=?;`, bookid)
	if err != nil {
		return err
	}

	// updating books
	_, err = database.Db.Exec(`UPDATE books
	SET IsDeleted=1
	WHERE BookId=? ;`, bookid)
	if err != nil {
		return err
	}
	return nil
}
