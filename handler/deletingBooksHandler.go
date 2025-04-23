package handler

import (
	"fmt"
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete Books
// @Description Delete a book
// @Tags Creating
// @Accept json
// @Produce json
// @Param id query string true "Book Identification"
// @Success 200 {object} DeleteBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /book [delete]
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
	sendSuccess(c, "book-deleted", creating)
}
