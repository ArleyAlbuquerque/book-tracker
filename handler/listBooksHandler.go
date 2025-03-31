package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

func ListBooksHandler(c *gin.Context) {
	books := []schemas.Creating{}

	if err := db.Find(&books).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(c, "list-books", books)

}
