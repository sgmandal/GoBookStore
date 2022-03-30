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
		v1.GET("/getgetbyauthor/:authorname", handler.GetGetAllByAuthor)
		v1.GET("/getbypublisher/:publishername", handler.GetAllByPublisher)
		v1.GET("/getbookbyid/:id", handler.GetBookById)
		v1.POST("/postabook", handler.PostABook)
		v1.PUT("/updateabook/:bookid", handler.UpdateABook)

		// purchasing a book using post request
		// getting book details

		v1.GET("/purchase/:bookid", handler.PurchaseABook)

		/*delete a book
		goal: its inventory should also be deleted if it exists
		do a soft delete

		approach:
		delete book by id
		need to do a softdelete so implement a tinyint books
		before printing book check whether isdeleted equals zero or not
		could either check book entry first to print inventory or just put isdeleted column on inventory too for now
		*/
		// end goal, remove isdeleted term from display
		v1.DELETE("/deleteabook/:bookid", handler.DeleteABook)
	}

	return router
}
