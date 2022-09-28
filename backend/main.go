package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/WibuSOS/sinarmas/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type DataRequest struct {
	Text string `json:"text"`
}

func handler(c *gin.Context) {
	db := database.GetDB()
	c.JSON(http.StatusOK, gin.H{
		"data": db.Data,
	})
}

func postHandler(c *gin.Context) {
	var data DataRequest
	db := database.GetDB()
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Append(data.Text)
	fmt.Println(db)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": data.Text})
}

func main() {
	database.StartDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/", handler)
	r.POST("/send", postHandler)

	r.Run(":" + os.Getenv("PORT"))
}
