package controllers

import (
	"errors"
	"net/http"
	"server/databases"
	"server/dtos"
	"server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var userReq dtos.UserRegisterRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.CreateUser(userReq.Username, userReq.Password, userReq.FullName)

	if err := databases.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "created",
		"data": gin.H{
			"user": newUser,
		},
	})
}

func Login(c *gin.Context) {
	var userReq dtos.UserLoginRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := databases.DB.Where("username = ?", userReq.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Incorrect username or password!"})
			// c.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "System error",
			"message": err,
		})
		return
	}

	if user.Password != userReq.Password {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "Incorrect username or password!",
			// "message": "Wrong password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": user,
		},
	})
}

func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": user,
		},
	})
}
