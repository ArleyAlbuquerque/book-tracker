package router

import (
	docs "github.com/ArleyAlbuquerque/book-tracker/docs"
	"github.com/ArleyAlbuquerque/book-tracker/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		//SHOW BOOKS
		v1.GET("/book", handler.ShowingBooksHandler)
		v1.POST("/create", handler.CreatingBooksHandler)
		v1.PUT("/update", handler.UpdatingBooksHandler)
		v1.DELETE("/delete", handler.DeletingBooksHandler)
		v1.GET("/showbooks", handler.ListBooksHandler)

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
