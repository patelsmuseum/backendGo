package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get data call api",
		})
	})

	router.POST("/getPost", func(c *gin.Context) {
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		c.JSON(200, gin.H{
			"message":  "Post data call api",
			"bodyData": string(value),
		})
	})

	router.Run(":5050")

}
