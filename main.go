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
	CommentController := &controllers.CommentDB{DB: db}

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/users/register", UserController.Register)
	router.POST("/users/login", UserController.Login)
	router.Use(middlewares.AuthJWT())
	// users
	router.PUT("/users/:id", UserController.UserUpdate)
	router.DELETE("/users", UserController.UserDelete)
	// photos
	router.POST("/photos", PhotoController.CreatedPhoto)
	router.GET("/photos", PhotoController.GetPhotos)
	router.PUT("/photos/:id", PhotoController.UpdatePhoto)
	router.DELETE("/photos/:id", PhotoController.DeletePhoto)
	// comments
	router.POST("/comments", CommentController.CreateComment)
	router.GET("/comments", CommentController.GetComments)

	router.Run(":3000")
}