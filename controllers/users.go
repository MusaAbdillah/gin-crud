package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gin-crud/models"
	"gin-crud/schemas"
	"net/http"
	"fmt"
)

func Register(ctx *gin.Context){
	fmt.Println("==Register==")
	
	var (
		input schemas.UserRegister
		user models.User
	)

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}  

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; !(gorm.IsRecordNotFoundError(err)) {
		ctx.JSON(401, gin.H{"error": "Email already taken!"})
		return
	}

	userRecord := models.User{Email: input.Email, Password: input.Password}
	models.DB.Create(&userRecord)

	ctx.JSON(200, gin.H{"data": "Register"})
}

func Index(ctx *gin.Context){
	fmt.Println("==Index==")

	var (
		users []models.User
	)

	models.DB.Find(&users)
	ctx.JSON(200, gin.H{"data": users})
}


func Show(ctx *gin.Context) {
	fmt.Println("==Show==")
	
	id := ctx.Param("id")

	var (
		user models.User
	)

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"data": user})
}