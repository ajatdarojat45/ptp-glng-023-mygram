package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type UserDB struct {
	DB *gorm.DB
}

func (db *UserDB) Register(c *gin.Context){
	c.JSON(201, gin.H{
		"age": "int",
		"email": "string",
		"id": "int",
		"username": "string",
	})
}

func (db *UserDB) Login(c *gin.Context){
	c.JSON(201, gin.H{
		"token": "string",
	})
}

func (db *UserDB) UserUpdate(c *gin.Context){
	c.JSON(201, gin.H{
		"id": "int",
		"age": "int",
		"email": "string",
		"username": "string",
	})
}

func (db *UserDB) UserDelete(c *gin.Context){
	c.JSON(201, gin.H{
		"message": "Your account has been successfully deleted",
	})
}