package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Show all books",
	})
}
