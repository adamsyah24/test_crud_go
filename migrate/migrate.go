package main

import (
	"github.com/adamsyah24/test-crud/initializers"
	"github.com/adamsyah24/test-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
