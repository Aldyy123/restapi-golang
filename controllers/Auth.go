package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}

func Register(c *gin.Context) {
	var user User

	if c.ShouldBindWith(&user, binding.Form) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field is required",
		})
		c.Abort()
		return
	}
	hash, err := HashingPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error hashing password",
		})
		c.Abort()
		return
	}

	user.Password = hash
	token, errToken := GenerateJWT(user.Username)

	if errToken != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "error generating token",
		})
	}

	var model = models.Users{
		Username: user.Username,
		Password: user.Password,
		Token:    token,
	}

	db := OpenSqlite()

	fmt.Println(&model)
	db.Create(&model)

	c.JSON(http.StatusOK, gin.H{
		"person": user,
		"token":  token,
	})

}

func Login(c *gin.Context) {
	var user User
	db := OpenSqlite()

	errJson := c.ShouldBindJSON(&user)

	if errJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field is required",
			"error":   errJson.Error(),
		})
		c.Abort()
		return
	}

	modelUsers := models.Users{
		Username: user.Username,
		Password: user.Password,
	}

	err := db.Where("username = ?", user.Username).First(&modelUsers).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	isValidPass := ValidationPassword(user.Password, modelUsers.Password)

	if isValidPass {

		token, errToken := GenerateJWT(user.Username)

		if errToken != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"message": "error generating token",
			})
		}

		modelUsers.Token = token

		db.Save(&modelUsers)

		c.JSON(http.StatusOK, gin.H{
			"user": modelUsers,
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Gagal password",
	})

}

func Authentication(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token is required",
		})
		c.Abort()
		return
	}
	splitBearer := strings.Split(authorization, " ")
	user, err := ValidationToken(splitBearer[1])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()

}
