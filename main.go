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
	photoRouter := router.Group("/photos")
	{
		photoRouter.POST("/", PhotoController.CreatedPhoto)
		photoRouter.GET("/", PhotoController.GetPhotos)
		photoRouter.Use(middlewares.AuthzPhoto())
		photoRouter.PUT("/:photoId", PhotoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", PhotoController.DeletePhoto)
	}
	// comments
	commentRouter := router.Group("/comments")
	{
		commentRouter.POST("/", CommentController.CreateComment)
		commentRouter.GET("/", CommentController.GetComments)
		commentRouter.Use(middlewares.AuthzComment())
		commentRouter.PUT("/:commentId", CommentController.UpdateComment)
		commentRouter.DELETE("/:commentId", CommentController.DeleteComment)
	}
	// social media
	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.POST("/", SocialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", SocialMediaController.GetSocialMedias)
		socialMediaRouter.Use(middlewares.AuthzSocialMedia())
		socialMediaRouter.PUT("/:socialMediaId", SocialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", SocialMediaController.DeleteSocialMedia)
	}

	router.Run(":3000")
}