package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type PhotoDB struct {
	DB *gorm.DB
}

func (db *PhotoDB) CreatedPhoto(c *gin.Context){
	c.JSON(201, gin.H{
		"id": "int",
		"title": "string",
		"caption": "int",
		"photo_url": "string",
		"user_id": "int",
		"created_at": "date",
	})
}

func (db *PhotoDB) GetPhotos(c *gin.Context){
	c.JSON(200, gin.H{
		"id": "int",
		"title": "string",
		"caption": "int",
		"photo_url": "string",
		"user_id": "int",
		"created_at": "date",
		"user": "obj",
	})
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