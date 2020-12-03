package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

const (

	host     = "localhost"
	port     = 5432
	user     = "sammaus"
	dbname   = "calculator"
)

func main(){
	var port_num string = "5000"
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Yes" : "I'm working",
		})
	})

	router.GET("/api/equations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name" : "sam",
		})
	})
	router.Run(":" + port_num)
	fmt.Printf("Listening on port 5000")
}