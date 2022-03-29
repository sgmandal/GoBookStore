package api

import (
	"bookstore/handler"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/getallbooks", handler.GetAllBooks)
		v1.POST("/getbyauthor", handler.GetAllByAuthor)

	}

	return router
}
