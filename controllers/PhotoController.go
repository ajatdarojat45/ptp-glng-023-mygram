package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"mygram/models"
	"net/http"
	"strconv"
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

	if req.Title == "" {
		c.JSON(400, gin.H{
			"message": "Title is required",
		})
		return
	}

	if req.Photo_Url == "" {
		c.JSON(400, gin.H{
			"message": "Photo_Url is required",
		})
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
	id := c.Param("photoId")
	photoId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "params photoId is required",
		})
		return
	}

	var photo models.Photo
	errPhoto := db.DB.First(&photo, photoId).Error
	if errPhoto != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	req := models.Photo{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Title == "" {
		c.JSON(400, gin.H{
			"message": "Title is required",
		})
		return
	}

	if req.Photo_Url == "" {
		c.JSON(400, gin.H{
			"message": "Photo_Url is required",
		})
		return
	}

	errUpdate := db.DB.Model(&photo).Updates(models.Photo{Title: req.Title, Caption: req.Caption, Photo_Url: req.Photo_Url}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": photo.ID,
		"title": photo.Title,
		"caption": photo.Caption,
		"photo_url": photo.Photo_Url,
		"user_id": photo.User_Id,
		"created_at": photo.Created_At,
	})
}

func (db *PhotoDB) DeletePhoto(c *gin.Context){
	var (
		photo models.Photo
	)

	id := c.Param("photoId")
	photoId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := db.DB.First(&photo, photoId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return 
	} else {
		errDelete := db.DB.Delete(&photo).Error
		if errDelete != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return 
		}else {
			c.JSON(200, gin.H{
				"message": "Your photo has been successfully deleted",
			})
			return 
		}
	}
}