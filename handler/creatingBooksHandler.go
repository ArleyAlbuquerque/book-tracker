package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatingBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Book created",
	})
}
