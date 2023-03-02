package main

import (
	"fmt"
	"gin-web-curd/initializers"
	"gin-web-curd/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my room",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		var body struct {
			Date    string `json:"date"`
			Content string `json:"content"`
			Author  string `json:"author"`
		}

		err := c.Bind(&body)
		if err != nil {
			return
		}
		log.Println(&body)
		data := models.Blog{
			Date:    body.Date,
			Content: body.Content,
			Author:  body.Author,
		}
		initializers.DB.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"message": "create record OK",
		})
	})

	r.GET("/list", func(c *gin.Context) {
		var data []models.Blog

		initializers.DB.Find(&data)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
