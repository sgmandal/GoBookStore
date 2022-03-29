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

func GetAllBooks(c *gin.Context) {
	fmt.Println("in GetAllBooks")
	var books []models.Book
	err := helper.GetEveryBook(&books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, books)
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

func GetGetAllByAuthor(c *gin.Context) {

	AuthorName := c.Params.ByName("authorname")

	// fetching all the books written by the writer
	var books []models.Book
	err := helper.GetByAuthor(AuthorName, &books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, books)
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
		c.JSON(http.StatusOK, books)
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
		c.JSON(http.StatusOK, books)
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
		c.JSON(http.StatusOK, x)
	}

}
