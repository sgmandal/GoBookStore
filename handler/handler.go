package handler

import (
	"bookstore/helper"
	"bookstore/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksWithout struct {
	Book_id     int    `json:"bookid"`
	AuthorID    int    `json:"AuthorID"`
	PublisherID int    `json:"PublisherID"`
	Book_Name   string `json:"BookName"`
	Book_Url    string `json:"BookURL"`
}

func GetAllBooks(c *gin.Context) {
	fmt.Println("in GetAllBooks")
	var books []models.Book
	err := helper.GetEveryBook(&books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		bookwithout := make([]BooksWithout, len(books))
		for i, k := range books {
			bookwithout[i].Book_id = k.Book_id
			bookwithout[i].AuthorID = k.AuthorID
			bookwithout[i].PublisherID = k.PublisherID
			bookwithout[i].Book_Name = k.Book_Name
			bookwithout[i].Book_Url = k.Book_Url
		}
		c.JSON(http.StatusOK, bookwithout)
	}
}

// func GetAllByAuthor(c *gin.Context) {

// 	// fetching authorname via post as it contains spaces
// 	var authorName models.Author
// 	c.BindJSON(&authorName)
// 	fmt.Println(authorName.AuthorName) // printing the same

// 	// fetching all the books written by the writer
// 	var books []models.Book
// 	err := helper.GetByAuthor(authorName.AuthorName, &books)
// 	if err != nil {
// 		c.AbortWithError(http.StatusNotFound, err)
// 	} else {
// 		c.JSON(http.StatusOK, books)
// 	}
// }

// struct {
// 	Book_id     int    `json:"bookid"`
// 	AuthorID    int    `json:"AuthorID"`
// 	PublisherID int    `json:"PublisherID"`
// 	Book_Name   string `json:"BookName"`
// 	Book_Url    string `json:"BookURL"`
// }{
// 	Book_id:     books[0].Book_id,
// 	PublisherID: books[0].PublisherID,
// 	Book_Name:   books[0].Book_Name,
// 	Book_Url:    books[0].Book_Url,
// }

func GetGetAllByAuthor(c *gin.Context) {

	AuthorName := c.Params.ByName("authorname")

	// fetching all the books written by the writer
	var books []models.Book

	err := helper.GetByAuthor(AuthorName, &books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		bookwithout := make([]BooksWithout, len(books))
		for i, k := range books {
			bookwithout[i].Book_id = k.Book_id
			bookwithout[i].AuthorID = k.AuthorID
			bookwithout[i].PublisherID = k.PublisherID
			bookwithout[i].Book_Name = k.Book_Name
			bookwithout[i].Book_Url = k.Book_Url
		}
		c.JSON(http.StatusOK, bookwithout)
	}
}

func GetAllByPublisher(c *gin.Context) {

	PublisherName := c.Params.ByName("publishername")

	// fetching all the books written by the writer
	var books []models.Book
	err := helper.GetByPublisher(PublisherName, &books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		bookwithout := make([]BooksWithout, len(books))
		for i, k := range books {
			bookwithout[i].Book_id = k.Book_id
			bookwithout[i].AuthorID = k.AuthorID
			bookwithout[i].PublisherID = k.PublisherID
			bookwithout[i].Book_Name = k.Book_Name
			bookwithout[i].Book_Url = k.Book_Url
		}
		c.JSON(http.StatusOK, bookwithout)
	}
}

func GetBookById(c *gin.Context) {
	bookid := c.Params.ByName("id")
	bookidint, err := strconv.Atoi(bookid)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	}

	var books []models.Book
	err = helper.GetBookByIDHelper(bookidint, &books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, struct {
			Book_id     int    `json:"bookid"`
			AuthorID    int    `json:"AuthorID"`
			PublisherID int    `json:"PublisherID"`
			Book_Name   string `json:"BookName"`
			Book_Url    string `json:"BookURL"`
		}{
			Book_id:     books[0].Book_id,
			PublisherID: books[0].PublisherID,
			Book_Name:   books[0].Book_Name,
			Book_Url:    books[0].Book_Url,
		})
	}
}

func PostABook(c *gin.Context) {
	var books models.Book
	c.BindJSON(&books)
	fmt.Println(books)

	err := helper.PostABookHandler(&books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, books)
	}
}

func UpdateABook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("bookid"))
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, errors.New("enter a vaild id"))
		return
	}

	var Book models.Book
	c.BindJSON(&Book)

	err = helper.UpdateABookHelper(id, &Book)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	} else {
		x := []models.Book{}
		helper.GetBookByIDHelper(id, &x)
		c.JSON(http.StatusOK, struct {
			Book_id     int    `json:"bookid"`
			AuthorID    int    `json:"AuthorID"`
			PublisherID int    `json:"PublisherID"`
			Book_Name   string `json:"BookName"`
			Book_Url    string `json:"BookURL"`
		}{
			Book_id:     x[0].Book_id,
			PublisherID: x[0].PublisherID,
			Book_Name:   x[0].Book_Name,
			Book_Url:    x[0].Book_Url,
		})
	}

}

func PurchaseABook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("bookid"))
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, errors.New("enter a vaild id"))
		return
	}
	err = helper.PurchaseABookHelper(id)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	} else {
		x := []models.Inventory{}
		helper.ShowInventoryHelper(&x)
		countx := make([]struct {
			BookId    int    `db:"BookId" json:"BookId"`
			NoOfBooks int    `db:"NoOfBooks" json:"NoOfBooks"`
			AddedBy   string `db:"AddedBy" json:"AddedBy"`
		}, len(x))
		for i, k := range x {
			countx[i].BookId = k.BookId
			countx[i].NoOfBooks = k.NoOfBooks
			countx[i].AddedBy = k.AddedBy

		}
		c.JSON(http.StatusOK, countx)
	}
}

func DeleteABook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("bookid"))
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, errors.New("enter a vaild id"))
		return
	}
	err = helper.DeleteABookHelper(id)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	} else {
		x := []models.Book{}
		helper.GetEveryBook(&x)
		c.JSON(http.StatusOK, x)
	}
}
