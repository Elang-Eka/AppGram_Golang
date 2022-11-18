package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Read User
func UserGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

// Read User
func UserGetAll(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := []models.User{}
	userID := uint(userData["id"].(float64))
	_ = userID
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}
	type GetUser struct {
		User_Id  int    `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	var data []GetUser
	for i := 0; i < len(User); i++ {
		StructData := GetUser{
			User_Id:  int(User[i].ID),
			Username: User[i].Username,
			Email:    User[i].Email,
		}
		data = append(data, StructData)
	}

	c.JSON(http.StatusOK, gin.H{
		"any_user": data,
	})
}

// Update User
func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Model(&User).Where("id = ?", userID).Updates(models.User{Username: User.Username, Email: User.Email, Age: User.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

// Delete User
func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Model(&User).Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&User).Where("id = ?", userID).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
