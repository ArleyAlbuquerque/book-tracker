package handler

import (
	"fmt"
	"net/http"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("operation from handler %s sucessful", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateBookResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CreatingResponses `json:"data"`
}

type DeleteBookResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CreatingResponses `json:"data"`
}
type ShowBookResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CreatingResponses `json:"data"`
}

type ListBookResponse struct {
	Message string                      `json:"message"`
	Data    []schemas.CreatingResponses `json:"data"`
}
type UpdateBookResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CreatingResponses `json:"data"`
}
