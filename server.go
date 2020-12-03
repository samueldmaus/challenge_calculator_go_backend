package main

import (
	"fmt"


	"github.com/gin-gonic/gin"
)

func main(){
	var port_num string = "5000"
	router := gin.New()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name" : "sam",
		})
	})
	router.Run(":" + port_num)
	fmt.Printf("Listening on port 5000")
}