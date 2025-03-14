package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

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

	sendSucess(c, "creating-book", creating)
}
