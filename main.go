package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Go Server!", nil)
	})

	router.Run(":" + port)
}
