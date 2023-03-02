package main

import (
	"gin-web-curd/initializers"
	"gin-web-curd/models"
)

func init() {
	initializers.ConnectDB()
}
func main() {
	err := initializers.DB.AutoMigrate(models.Blog{})
	if err != nil {
		return
	}
}
