package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Books
// @Description show all books
// @Tags Creating
// @Accept json
// @Produce json
// @Param  id query string true "Book Identification"
// @Success 200 {object} ShowBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /book [get]
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

	sendSuccess(c, "show-book", creating)
}
