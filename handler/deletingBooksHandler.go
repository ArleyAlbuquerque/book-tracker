package handler

import (
	"fmt"
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

func DeletingBooksHandler(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	creating := schemas.Creating{}

	if err := db.First(&creating, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", id))
		return
	}

	if err := db.Delete(&creating).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting book with id: %s", id))
		return
	}
	sendSucess(c, "book-deleted", creating)
}
