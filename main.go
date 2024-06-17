package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-yields-api/models"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Initialize GORM with SQLite (or any other DB of your choice)
	db, err := gorm.Open(sqlite.Open("yields.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Yield{})

	// Define routes
	r.GET("/yields", func(c *gin.Context) {
		client := resty.New()
		resp, err := client.R().Get("https://yields.llama.fi/pools")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var yields []models.Yield
		err = json.Unmarshal(resp.Body(), &yields)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Save data to the database (optional)
		for _, yield := range yields {
			db.Create(&yield)
		}

		c.JSON(http.StatusOK, yields)
	})

	// Run the server
	r.Run(":8080")
}
