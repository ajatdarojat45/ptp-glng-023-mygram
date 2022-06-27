package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"mygram/models"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"mygram/helpers"
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
	var req models.User

	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err);
	}

	dbResult := models.User{}
	errUser := db.DB.Debug().Where("email = ?", req.Email).Last(&dbResult).Error
	if errUser != nil {
		c.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(dbResult.Password), []byte(req.Password))
	if errBcrypt != nil {
		c.AbortWithError(http.StatusBadRequest, errBcrypt)
		return
	}

	token := helpers.GenerateToken(req.Username)

	c.JSON(200, gin.H{
		"token": token,
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