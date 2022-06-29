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