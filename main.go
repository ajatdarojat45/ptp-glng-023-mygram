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
	SocialMediaController := &controllers.SocialMediaDB{DB: db}

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/users/register", UserController.Register)
	router.POST("/users/login", UserController.Login)
	router.Use(middlewares.AuthJWT())
	// users
	router.PUT("/users/:userId", UserController.UserUpdate)
	router.DELETE("/users", UserController.UserDelete)
	// photos
	router.POST("/photos", PhotoController.CreatedPhoto)
	router.GET("/photos", PhotoController.GetPhotos)
	router.PUT("/photos/:photoId", PhotoController.UpdatePhoto)
	router.DELETE("/photos/:photoId", PhotoController.DeletePhoto)
	// comments
	router.POST("/comments", CommentController.CreateComment)
	router.GET("/comments", CommentController.GetComments)
	router.PUT("/comments/:commentId", CommentController.UpdateComment)
	router.DELETE("/comments/:commentId", CommentController.DeleteComment)
	// social media
	router.POST("/socialmedias", SocialMediaController.CreateSocialMedia)
	router.GET("/socialmedias", SocialMediaController.GetSocialMedias)
	router.PUT("/socialmedias/:socialMediaId", SocialMediaController.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", SocialMediaController.DeleteSocialMedia)

	router.Run(":3000")
}