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
		//v1.POST("/getbyauthor", handler.GetAllByAuthor)
		//v1.POST("getbypublisher", handler.GetAllByPublisher)
		v1.GET("/getgetbyauthor/:authorname", handler.GetGetAllByAuthor)
		v1.GET("/getbypublisher/:publishername", handler.GetAllByPublisher)
		v1.GET("/getbookbyid/:id", handler.GetBookById)
		v1.POST("/postabook", handler.PostABook)
		v1.PUT("/updateabook/:bookid", handler.UpdateABook)
	}

	return router
}
