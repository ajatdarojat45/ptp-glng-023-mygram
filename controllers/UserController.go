package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"mygram/models"
	"net/http"
)

type UserDB struct {
	DB *gorm.DB
}

func (db *UserDB) Register(c *gin.Context){
	var req models.User

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
		"age": req.Age,
		"email": req.Email,
		"id": req.ID,
		"username": req.Username,
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