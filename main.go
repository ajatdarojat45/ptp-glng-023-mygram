package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mygram/config"
	"mygram/controllers"
	"mygram/middlewares"
)

func main(){
	router := gin.Default()
	db := config.DBConnect()
	UserController := &controllers.UserDB{DB: db}
	PhotoController := &controllers.PhotoDB{DB: db}

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/users/register", UserController.Register)
	router.POST("/users/login", UserController.Login)
	router.Use(middlewares.AuthJWT())
	router.PUT("/users", UserController.UserUpdate)
	router.POST("/photos", PhotoController.CreatedPhoto)

	router.Run(":3000")
}