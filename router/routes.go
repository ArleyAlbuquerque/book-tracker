package router

import (
	"github.com/ArleyAlbuquerque/book-tracker/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//SHOW BOOKS
		v1.GET("/book", handler.ShowingBooksHandler)
		v1.POST("/create", handler.CreatingBooksHandler)
		v1.PUT("/update", handler.UpdatingBooksHandler)
		v1.DELETE("/delete", handler.DeletingBooksHandler)
		v1.GET("/showbooks", handler.ListBooksHandler)

	}
}
