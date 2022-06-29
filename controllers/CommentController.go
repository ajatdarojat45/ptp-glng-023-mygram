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
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Message == "" {
		c.JSON(400, gin.H{
			"message": "Message is required",
		})
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

func (db *CommentDB) UpdateComment(c *gin.Context){
	id := c.Param("commentId")
	commentId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "params photoId is required",
		})
		return
	}

	var comment models.Comment
	errComment := db.DB.First(&comment, commentId).Error
	if errComment != nil {
		c.JSON(400, gin.H{
			"message": "Data not found",
		})
		return
	}

	req := models.Comment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Message == "" {
		c.JSON(400, gin.H{
			"message": "Message is required",
		})
		return
	}

	errUpdate := db.DB.Model(&comment).Updates(models.Comment{Message: req.Message}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(201, gin.H{
		"id": comment.ID,
		"message": comment.Message,
		"photo_id": comment.Photo_Id,
		"user_id": comment.User_Id,
		"created_at": comment.Created_At,
	})
}

func (db *CommentDB) DeleteComment(c *gin.Context){
	var (
		comment models.Comment
	)

	id := c.Param("commentId")
	commentId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := db.DB.First(&comment, commentId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return 
	} else {
		errDelete := db.DB.Delete(&comment).Error
		if errDelete != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return 
		}else {
			c.JSON(200, gin.H{
				"message": "Your comment has been successfully deleted",
			})
			return 
		}
	}
}