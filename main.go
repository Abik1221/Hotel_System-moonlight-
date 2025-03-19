package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hellow world", "name": "namhom Keneni"}) })

	router.Run(":" + port)

}
