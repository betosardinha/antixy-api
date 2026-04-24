package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/betosardinha/antixy-api/internal/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	_ = database

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Antixy API running",
		})
	})

	r.Run(":8080")
}
