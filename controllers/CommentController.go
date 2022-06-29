package controllers

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"mygram/models"
)

type CommentDB struct {
	DB *gorm.DB
}

func (db *CommentDB) CreateComment(c *gin.Context){
	var req models.Comment
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
		"message": req.Message,
		"photo_id": req.Photo_Id,
		"user_id": req.User_Id,
		"created_at": req.Created_At,
	})
}