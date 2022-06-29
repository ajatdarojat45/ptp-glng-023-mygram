package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"mygram/models"
	"net/http"
	"strconv"
	// "fmt"
)

type PhotoDB struct {
	DB *gorm.DB
}

func (db *PhotoDB) CreatedPhoto(c *gin.Context){
	var req models.Photo
	userId := c.GetString("userId")
	userIdConvert, _ := strconv.Atoi(userId)
	req.User_Id = userIdConvert

	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	errCreate := db.DB.Debug().Create(&req).Error
	if errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(201, gin.H{
		"id": req.ID,
		"title": req.Title,
		"caption": req.Caption,
		"photo_url": req.Photo_Url,
		"user_id": req.User_Id,
		"created_at": req.Created_At,
	})
}

func (db *PhotoDB) GetPhotos(c *gin.Context){
	var (
		photos []models.Photo
		photosRes []models.PhotoList
	)

	db.DB.Preload("User").Find(&photos)
	for _,el := range photos{
		photosRes = append(photosRes, models.PhotoList{
			ID: el.ID,
			Title: el.Title,
			Caption: el.Caption,
			Photo_Url: el.Photo_Url,
			User_Id: el.User_Id,
			Created_At: el.Created_At,
			Updated_At: el.Updated_At,
			User: models.PhotoListUser{
				Email: el.User.Email,
				Username: el.User.Username,
			},
		})
	}

	c.JSON(200, photosRes)
}

func (db *PhotoDB) UpdatePhoto(c *gin.Context){
	c.JSON(201, gin.H{
		"id": "int",
		"title": "string",
		"caption": "int",
		"photo_url": "string",
		"user_id": "int",
		"created_at": "date",
	})
}

func (db *UserDB) DeletePhoto(c *gin.Context){
	c.JSON(201, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}