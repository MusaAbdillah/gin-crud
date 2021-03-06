package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
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


// FindBooks godoc
// @Summary FindBooks
// @Description fetch list of book
// @ID FindBooks
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Book
// @Header 200 {string} Token "qwerty"
// @Router /books [get]
func FindBooks(c *gin.Context) {
	fmt.Println("==Controller F`indBooks==")
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook godoc
// @Summary CreateBook
// @Description Crete book record
// @ID CreateBook
// @Accept  json
// @Produce  json
// @Param title path string true "Title"
// @Param author path string true "Author"
// @Success 200 {object} models.Book
// @Header 200 {string} Token "qwerty"
// @Router /books [post]
func CreateBook(c *gin.Context) {
	fmt.Println("==Controller CreateBook==")
	var input CreateBookInput
	var bookModel models.Book
	
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}  

	if err := models.DB.Where("title = ?", input.Title).First(&bookModel).Error; !(gorm.IsRecordNotFoundError(err)) {
		c.JSON(401, gin.H{"error": "Title already taken!"})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBook godoc
// @Summary FindBook
// @Description Find specify book
// @ID FindBook
// @Accept  json
// @Produce  json
// @Param title path string true "Title"
// @Success 200 {object} models.Book
// @Header 200 {string} Token "qwerty"
// @Router /books/:title [get]
func FindBook(c *gin.Context) {
	fmt.Println("==Controller FindBook==")
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary UpdateBook
// @Description Update specify book
// @ID UpdateBook
// @Accept  json
// @Produce  json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Success 200 {object} models.Book
// @Header 200 {string} Token "qwerty"
// @Router /books/:title [patch]
func UpdateBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	input_err := c.ShouldBindJSON(&input)

	if input_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary DeleteBook
// @Description Delete specify book
// @ID DeleteBook
// @Accept  json
// @Produce  json
// @Param title path string true "Title"
// @Success 200 {object} models.Book
// @Header 200 {string} Token "qwerty"
// @Router /books/{:title} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&book).Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted!"})

}