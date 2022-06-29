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

func (db *CommentDB) GetComments(c *gin.Context){
	var (
		comments []models.Comment
		commentsRes []models.CommentList
	)

	db.DB.Preload("User").Preload("Photo").Find(&comments)
	for _,el := range comments{
		commentsRes = append(commentsRes, models.CommentList{
			ID: el.ID,
			Message: el.Message,
			Photo_Id: el.Photo_Id,
			User_Id: el.User_Id,
			Created_At: el.Created_At,
			Updated_At: el.Updated_At,
			User: models.CommentListUser{
				ID: el.User.ID,
				Email: el.User.Email,
				Username: el.User.Username,
			},
			Photo: models.CommentListPhoto{
				ID: el.Photo.ID,
				Title: el.Photo.Title,
				Caption: el.Photo.Caption,
				Photo_Url: el.Photo.Photo_Url,
				User_Id: el.Photo.User_Id,
			},
		})
	}

	c.JSON(200, commentsRes)
}