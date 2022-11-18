package controllers

import (
	"log"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DataP struct {
	ID        uint      `json:"photo_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	User      DataP2
}

type DataP2 struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// create photo
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))

	file, err := c.FormFile("photo_url")
	if err != nil {
		log.Fatal(err)
	}

	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_Id = userID
	Photo.Photo_url = file.Filename

	idPhoto := strconv.Itoa(int(userID))
	createTimeM := strconv.Itoa(Photo.CreatedAt.Minute())
	createTimeD := strconv.Itoa(Photo.CreatedAt.Day())
	createTimeH := strconv.Itoa(Photo.CreatedAt.Hour())
	location := "assets/img/" + idPhoto + createTimeM + createTimeH + createTimeD + file.Filename
	Photo.Photo_url = location
	err = c.SaveUploadedFile(file, location)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               Photo.ID,
		"title":            Photo.Title,
		"social_media_url": Photo.Caption,
		"photo_url":        location,
		"user_id":          userID,
		"created_at":       Photo.CreatedAt,
	})
}

// Read User
func PhotoGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := []models.Photo{}
	User := models.User{}

	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("user_id = ?", userID).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	err = db.Select("username, id").Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
	}

	users := DataP2{
		ID:       int(userID),
		Username: User.Username,
	}
	var data []DataP
	for i := 0; i < len(Photo); i++ {
		dataStruct := DataP{
			ID:        Photo[i].ID,
			Title:     Photo[i].Title,
			Caption:   Photo[i].Caption,
			Photo_url: Photo[i].Photo_url,
			UserID:    int(Photo[i].User_Id),
			CreatedAt: Photo[i].CreatedAt,
			UpdateAt:  Photo[i].UpdatedAt,
			User:      users,
		}
		data = append(data, dataStruct)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Update Photo
func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	file, err := c.FormFile("photo_url")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)
	err = c.SaveUploadedFile(file, "assets/img/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_Id = userID
	Photo.ID = uint(photoId)
	Photo.Photo_url = file.Filename

	err = db.Model(&Photo).Where("id = ?", photoId).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != Photo.User_Id {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, Photo_url: Photo.Photo_url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         photoId,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.User_Id,
		"updated_at": Photo.UpdatedAt,
	})
}

// deleted Photo
func PhotoDeleted(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_Id = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if userID != Photo.User_Id {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Delete(&Photo).Error
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
