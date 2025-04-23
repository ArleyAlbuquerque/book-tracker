package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Books
// @Description Create a new book
// @Tags Creating
// @Accept json
// @Produce json
// @Param request body CreatingBookRequest true "Request body"
// @Success 200 {object} CreateBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book [post]
func CreatingBooksHandler(c *gin.Context) {
	request := CreatingBookRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	creating := schemas.Creating{
		Name:  request.Name,
		Autor: request.Autor,
		Pages: request.Pages,
	}

	if err := db.Create(&creating).Error; err != nil {
		logger.ErrF("error creating book:%v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating on database")
		return
	}

	sendSuccess(c, "creating-book", creating)
}
