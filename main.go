package main

import (
	. "rest-api/controllers"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db := OpenSqlite()

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Books{})

	r := gin.Default()

	api := r.Group("/api")

	api.POST("/register", Register)
	api.POST("/login", Login)

	resource := r.Group("/resource").Use(Authentication)
	{
		resource.GET("/", GetBooks)
		resource.POST("/create", InsertBook)
		resource.GET("/:id", GetBookOne)
		resource.DELETE("/:id", DeleteBooks)
		resource.PUT("/:id", UpdateBooks)
	}

	r.Run()
}
