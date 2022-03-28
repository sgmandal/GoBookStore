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
	books, err := helper.GetEveryBook()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, books)
	}
}
