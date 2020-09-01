package controllers

import (
  "github.com/gin-gonic/gin"
  "gin-crud/models"
  "net/http"
  "fmt"
)

// schema 

type CreateBookInput struct {
	Title 	string `json:"title" 	binding:"required"`
	Author 	string `json:"author" 	binding:"required"`
}

type UpdateBookInput struct {
	Title 	string `json:"title"`
	Author 	string `json:"author`
}

func FindBooks(c *gin.Context) {
	fmt.Println("==Controller FindBooks==")
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	fmt.Println("==Controller CreateBook==")
	var input CreateBookInput
	
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}  

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	fmt.Println("==Controller FindBook==")
	var book models.Book
	err := models.DB.Where("title = ?", c.Param("title")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("title = ?", c.Param("title")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}