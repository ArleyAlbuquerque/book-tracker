package router

import "github.com/gin-gonic/gin"

func Initialize() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})
	r.Run(":8080")

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.Run(":8080")

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run(":8080")

	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.Run(":8080")

}
