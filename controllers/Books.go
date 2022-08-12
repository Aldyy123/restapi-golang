package controllers

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksForm struct {
	Title       string `form:"title" json:"title" binding:"required,min=3"`
	Description string `form:"description" json:"description" binding:"required,min=10"`
}

func GetBooks(c *gin.Context) {
	db := OpenSqlite()
	var books []models.Books

	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success tetrieved books",
		"books":   books,
	})
}

func InsertBook(c *gin.Context) {
	db := OpenSqlite()
	var booksForm BooksForm
	err := c.ShouldBindJSON(&booksForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	var books = models.Books{
		Title:       booksForm.Title,
		Description: booksForm.Description,
	}

	db.Create(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "Books created successfully",
		"books":   books,
	})

}

func GetBookOne(c *gin.Context) {
	paramID := c.Param("id")
	db := OpenSqlite()
	id, _ := strconv.Atoi(paramID)
	var books = models.Books{
		ID: id,
	}

	db.First(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"books":   books,
	})

}

func DeleteBooks(c *gin.Context) {
	paramID := c.Param("id")
	db := OpenSqlite()
	id, err := strconv.Atoi(paramID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Id must be a number",
		})
	}

	var books = models.Books{
		ID: id,
	}

	db.Delete(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete books successfully",
	})
}

func UpdateBooks(c *gin.Context) {
	paramID := c.Param("id")
	db := OpenSqlite()
	id, err := strconv.Atoi(paramID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Id must be a number",
		})
	}

	var books = models.Books{
		ID: id,
	}

	result := db.First(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusNotFound,
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	var booksForm BooksForm
	errJson := c.ShouldBindJSON(&booksForm)

	if errJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errJson,
		})
		c.Abort()
		return
	}

	books = models.Books{
		ID:          id,
		Title:       booksForm.Title,
		Description: booksForm.Description,
	}
	fmt.Println(books)
	resultUpdate := db.Save(&books)

	if resultUpdate.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusNotFound,
			"message": resultUpdate.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete books successfully",
	})
}
