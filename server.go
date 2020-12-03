package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

type Equations struct {
	ID int
	EQUATION string
	ANSWER string
}

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
	
	sqlStatement := `SELECT * FROM challenge_calculator LIMIT 10;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	  }
	  defer rows.Close()
	  
	  eqs := make([]Equations, 0)
	  for rows.Next() {
		eq := Equations{}
		err = rows.Scan(&eq.ID, &eq.EQUATION, &eq.ANSWER)
		if err != nil {
		  panic(err)
		}
		eqs = append(eqs, eq)
	  }
	  err = rows.Err()
	  if err != nil {
		panic(err)
	  }
	  fmt.Println(eqs)
	  c.JSON(200, eqs)
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