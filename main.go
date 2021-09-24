package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	engine := gin.Default()

	connectionString := "user=postgres host=database dbname=bio_consc port=54320" +
		" password=password sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		log.Println(err.Error())
		return
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	if err := engine.Run(); err != nil {
		log.Println(err.Error())
	}
}
