package main

import (
	"log"
	"os"

	"github.com/asapkau/hotdrop-lite/internal/cache"
	"github.com/asapkau/hotdrop-lite/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ no .env file found")
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	if err := cache.Init(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
