package controllers

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"mygram/models"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"mygram/helpers"
	"fmt"
	"strconv"
	"net/mail"
)

type UserDB struct {
	DB *gorm.DB
}

func (db *UserDB) Register(c *gin.Context){
	var (
		req models.User
		findUser models.User
	)

	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}

	if req.Username == "" {
		c.JSON(400, gin.H{
			"message": "Username is required",
		})
		return
	}

	if req.Password == "" {
		c.JSON(400, gin.H{
			"message": "Password is required",
		})
		return
	}

	if len(req.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "Minimun length of password is 6 char",
		})
		return
	}

	if req.Age < 8 {
		c.JSON(400, gin.H{
			"message": "Minimun age is 8 years",
		})
		return
	}

	_, errMailFormat := mail.ParseAddress(req.Email)
	if errMailFormat != nil {
		c.JSON(400, gin.H{
			"message": "Email format is warong",
		})
		return
	}

	db.DB.Where("email = ?", req.Email).First(&findUser)
	if findUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	db.DB.Where("username = ?", req.Username).First(&findUser)
	if findUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Username already used",
		})
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

	token := helpers.GenerateToken(dbResult.Username)

	c.JSON(200, gin.H{
		"token": token,
	})
}

func (db *UserDB) UserUpdate(c *gin.Context){
	id := c.Param("userId")
	userId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		fmt.Println("error found: ", errConvert)
		c.JSON(400, gin.H{
			"result": "params userId is required",
		})
		return
	}

	var user models.User
	var findUser models.User
	errUser := db.DB.First(&user, userId).Error
	if errUser != nil {
		c.JSON(404, gin.H{
			"result": "Data not found",
		})
		return
	}

	req := models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"result": "Bad Request",
		})
		return
	}

	if req.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}

	if req.Username == "" {
		c.JSON(400, gin.H{
			"message": "Username is required",
		})
		return
	}

	_, errMailFormat := mail.ParseAddress(req.Email)
	if errMailFormat != nil {
		c.JSON(400, gin.H{
			"message": "Email format is warong",
		})
		return
	}

	db.DB.Where("email = ?", req.Email).First(&findUser)
	if findUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	db.DB.Where("username = ?", req.Username).First(&findUser)
	if findUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Username already used",
		})
		return
	}

	errUpdate := db.DB.Model(&user).Updates(models.User{Username: req.Username, Email: req.Email}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"result": "internal server error",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"age": user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func (db *UserDB) UserDelete(c *gin.Context){
	var (
		user models.User
		result gin.H
	)

	id := c.GetString("userId");
	err := db.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		errDelete := db.DB.Delete(&user).Error
		if errDelete != nil {
			result = gin.H{
				"result": "delete failed",
			}
		}else {
			result = gin.H{
				"message": "Your account has been successfully deleted",
			}
		}
	}

	c.JSON(200, result)
}