package controllers

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"mygram/models"
)

type SocialMediaDB struct {
	DB *gorm.DB
}

func (db *SocialMediaDB) CreateSocialMedia(c *gin.Context){
	var req models.SocialMedia
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
		"name": req.Name,
		"social_media_url": req.Social_Media_Url,
		"user_id": req.User_Id,
		"created_at": req.Created_At,
	})
}

func (db *SocialMediaDB) GetSocialMedias(c *gin.Context){
	var (
		socialMedias []models.SocialMedia
		socialMediasRes []models.SocialMediaList
	)

	db.DB.Preload("User").Find(&socialMedias)

	for _,el := range socialMedias{
		socialMediasRes = append(socialMediasRes, models.SocialMediaList{
			ID: el.ID,
			Name: el.Name,
			Social_Media_Url: el.Social_Media_Url,
			User_Id: el.User_Id,
			Created_At: el.Created_At,
			Updated_At: el.Updated_At,
			User: models.SocialMediaListUser{
				ID: el.User.ID,
				Email: el.User.Email,
				Username: el.User.Username,
			},
		})
	}

	c.JSON(200, gin.H{
		"social_medias": socialMediasRes,
	})
}