package main

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"message": "Hello, World! from custom handler",
		"user":    name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)

	// router.Run()
	router.Run(":3000")
}
