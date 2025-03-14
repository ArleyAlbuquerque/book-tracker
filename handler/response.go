package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendSucess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("operation from handler %s sucessful", op),
		"data": data,
	})
}
