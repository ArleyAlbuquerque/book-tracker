package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Books
// @Description List all books
// @Tags Creating
// @Accept json
// @Produce json
// @Success 200 {object} ListBookResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func ListBooksHandler(c *gin.Context) {
	books := []schemas.Creating{}

	if err := db.Find(&books).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(c, "list-books", books)

}
