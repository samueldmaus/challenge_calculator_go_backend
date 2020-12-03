package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sammaus"
	dbname   = "calculator"
)

func sayHi(c *gin.Context) {
	c.JSON(200, gin.H{
		"I'm" : "working",
	})
}

func getEquations(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)
  	db, err := sql.Open("postgres", psqlInfo)
  	if err != nil {
    	panic(err)
  	}
  	defer db.Close()

  	err = db.Ping()
  	if err != nil {
    	panic(err)
  	}
	var id int
	var equation string
	var answer string
	sqlStatement := `SELECT * FROM challenge_calculator WHERE id=$1;`
	row := db.QueryRow(sqlStatement, 3)
	switch err := row.Scan(&id, &equation, &answer); err {
	case sql.ErrNoRows:
  		fmt.Println("No rows were returned!")
	case nil:
  		fmt.Println(id, equation, answer)
	default:
		  panic(err)
	}	
}

func main(){
	var port_num string = "5000"
	router := gin.New()

	router.GET("/", sayHi)

	router.GET("/api/equations", getEquations)

	router.Use(cors.Default())
	router.Run(":" + port_num)
	fmt.Printf("Listening on port 5000")
}