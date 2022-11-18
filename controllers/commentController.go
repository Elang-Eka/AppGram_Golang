package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// create comment
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	_ = userID

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Photo_id, _ := strconv.Atoi(c.Request.FormValue("photo_id"))

	err := db.Model(&Photo).Where("id = ?", Photo_id).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		Comment.User_id = userID
		Comment.Photo_id = uint(Photo_id)
		err = db.Debug().Create(&Comment).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":         Comment.ID,
			"message":    Comment.Message,
			"photo_id":   Comment.Photo_id,
			"user_id":    Comment.User_id,
			"created_at": Comment.CreatedAt,
		})
	}
}

// Get Comment
func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	User := []models.User{}
	Comment := []models.Comment{}
	userID := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	_ = userID

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Comment).Where("photo_id = ?", photoId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found Comment",
			"message": err.Error(),
		})
		return
	} else {
		type structComment1 struct {
			ID       uint   `json:"user_id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}
		type structComment2 struct {
			ID        uint   `json:"photo_id"`
			Title     string `json:"title"`
			Caption   string `json:"caption"`
			Photo_Url string `json:"photo_url"`
			User_Id   uint   `json:"user_id"`
		}
		type structComment3 struct {
			ID        uint      `json:"comment_id"`
			Message   string    `json:"message"`
			Photo_Id  uint      `json:"photo_id"`
			User_Id   uint      `json:"user_id"`
			CreatedAt time.Time `json:"created_at"`
			UpdateAt  time.Time `json:"update_at"`
			User      structComment1
			Photo     structComment2
		}

		var data []structComment3
		for i, v := range Comment {
			_ = v
			err := db.Where("id = ?", photoId).Find(&Photo).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found Photo",
					"message": err.Error(),
				})
			}

			err = db.Select("username, id, email").Where("id = ?", Comment[i].User_id).Find(&User).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found User",
					"message": err.Error(),
				})
			} else {
				var users structComment1
				for j, v := range User {
					_ = v
					users = structComment1{
						ID:       User[j].ID,
						Username: User[j].Username,
						Email:    User[j].Email,
					}
				}
				photos := structComment2{
					ID:        Photo.ID,
					Title:     Photo.Title,
					Caption:   Photo.Caption,
					Photo_Url: Photo.Photo_url,
					User_Id:   Photo.User_Id,
				}
				dataStruct := structComment3{
					ID:        Comment[i].ID,
					Message:   Comment[i].Message,
					Photo_Id:  Comment[i].Photo_id,
					User_Id:   Comment[i].User_id,
					UpdateAt:  Comment[i].UpdatedAt,
					CreatedAt: Comment[i].CreatedAt,
					User:      users,
					Photo:     photos,
				}
				data = append(data, dataStruct)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// update comment
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	User := models.User{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&User).Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != User.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	err = db.Model(&Comment).Where("id = ?", commentId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	Comment.User_id = userID

	c.JSON(http.StatusOK, gin.H{
		"id":         commentId,
		"message":    Comment.Message,
		"user_id":    Comment.User_id,
		"photo_id":   Comment.Photo_id,
		"updated_at": Comment.UpdatedAt,
	})
}

// deleted Comment
func CommentDeleted(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.User_id = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != Comment.User_id {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Another User",
			"message": "Data Cannot Be Deleted",
		})
		return
	}

	err = db.Model(&Comment).Where("id = ?", commentId).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been successfully deleted",
	})
}
