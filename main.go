package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	if err := engine.Run(); err != nil {
		log.Println(err.Error())
	}
}
