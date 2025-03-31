package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

func ShowingBooksHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	creating := schemas.Creating{}
	if err := db.First(&creating, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "book not found")
		return
	}

	sendSucess(c, "show-book", creating)
}
