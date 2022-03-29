package handler

import (
	"bookstore/helper"
	"bookstore/models"
	"fmt"
	"net/http"

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

func GetAllByAuthor(c *gin.Context) {

	// fetching authorname via post as it contains spaces
	var authorName models.Author
	c.BindJSON(&authorName)
	fmt.Println(authorName.AuthorName) // printing the same

	// fetching all the books written by the writer
	var books []models.Book
	err := helper.GetByAuthor(authorName.AuthorName, &books)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, books)
	}
}
