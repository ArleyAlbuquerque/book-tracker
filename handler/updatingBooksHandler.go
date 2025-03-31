package handler

import (
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

func UpdatingBooksHandler(c *gin.Context) {
	request := UpdateBookRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error:%v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

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
	//Update book
	if request.Name != "" {
		creating.Name = request.Name
	}
	if request.Autor != "" {
		creating.Autor = request.Autor
	}
	if request.Pages > 0 {
		creating.Pages = request.Pages
	}

	//save
	if err := db.Save(&creating).Error; err != nil {
		logger.ErrF("error updating book: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error updating book")
		return
	}

	sendSucess(c, "uptade-book", creating)

}
