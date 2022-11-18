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

type Data struct {
	ID               uint      `json:"sosmed_id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	UserID           int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdateAt         time.Time `json:"update_at"`
	User             Data2
}

type Data2 struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func CreateSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Sosmed := models.SocialMed{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserId = userID

	err := db.Debug().Create(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.Social_media_url,
		"user_id":          userID,
		"created_at":       Sosmed.CreatedAt,
	})
}

// Read User
func SosmedGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := []models.SocialMed{}
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	err := db.Where("user_id = ?", userID).Find(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}
	var data []Data
	for i := 0; i < len(Sosmed); i++ {
		err = db.Select("username, id").Where("id = ?", userID).Find(&User).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})
		}

		users := Data2{
			ID:       int(userID),
			Username: User.Username,
		}

		dataStruct := Data{
			ID:               Sosmed[i].ID,
			Name:             Sosmed[i].Name,
			Social_media_url: Sosmed[i].Social_media_url,
			UserID:           int(Sosmed[i].UserId),
			CreatedAt:        Sosmed[i].CreatedAt,
			UpdateAt:         Sosmed[i].UpdatedAt,
			User:             users,
		}
		data = append(data, dataStruct)
	}
	c.JSON(http.StatusOK, gin.H{
		"social_medias": data,
	})
}

// Update sosmed
func SosmedUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMed{}

	sosmedId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserId = userID
	Sosmed.ID = uint(sosmedId)

	if userID != Sosmed.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err := db.Model(&Sosmed).Where("id = ?", sosmedId).Updates(models.SocialMed{Name: Sosmed.Name, Social_media_url: Sosmed.Social_media_url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               sosmedId,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.Social_media_url,
		"user_id":          Sosmed.UserId,
		"updated_at":       Sosmed.UpdatedAt,
	})
}

// deleted sosmed
func SosmedDeleted(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMed{}

	sosmedId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserId = userID
	Sosmed.ID = uint(sosmedId)

	err := db.Model(&Sosmed).Where("id = ?", sosmedId).Find(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if userID != Sosmed.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad Request",
			"message": "Not Found Data Id",
		})
		return
	}

	err = db.Model(&Sosmed).Where("id = ?", sosmedId).Delete(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}

func SosmedGetByIdUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := []models.SocialMed{}
	_ = userData
	userId, _ := strconv.Atoi(c.Param("userIdSos"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	err := db.Model(&Sosmed).Where("user_id = ?", userId).Find(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	type sosmeds struct {
		ID               uint      `json:"sosmed_id"`
		Name             string    `json:"name"`
		Social_media_url string    `json:"social_media_url"`
		UserID           int       `json:"user_id"`
		CreatedAt        time.Time `json:"created_at"`
		UpdateAt         time.Time `json:"update_at"`
	}
	var data []sosmeds
	for i := 0; i < len(Sosmed); i++ {
		dataStruct := sosmeds{
			ID:               Sosmed[i].ID,
			Name:             Sosmed[i].Name,
			Social_media_url: Sosmed[i].Social_media_url,
			UserID:           int(Sosmed[i].UserId),
			CreatedAt:        Sosmed[i].CreatedAt,
			UpdateAt:         Sosmed[i].UpdatedAt,
		}
		data = append(data, dataStruct)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
