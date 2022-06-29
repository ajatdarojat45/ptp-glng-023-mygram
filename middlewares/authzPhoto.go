package middlewares

import (
	"errors"
	"net/http"
	"mygram/config"
	"mygram/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AuthzPhoto() gin.HandlerFunc {
	return func(c *gin.Context){
		id := c.Param("photoId")
		idConvert, errConvert := strconv.Atoi(id)
		if errConvert != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Bad request"))
			c.JSON(400, gin.H{
				"message": "params photoId is required",
			})
			return
		}

		result := models.Photo{}
		errFind := config.DBConnect().First(&result, idConvert).Error
		if errFind != nil {
			c.AbortWithError(404, errors.New("Data not found"))
			c.JSON(404, gin.H{
				"message": "Data not found",
			})
			return
		}else {
			userId := c.GetString("userId")
			userIdConvert, _ := strconv.Atoi(userId)
			if result.User_Id != userIdConvert {
				c.AbortWithError(403, errors.New("Forbidden access"))
				c.JSON(404, gin.H{
					"message": "Forbidden access",
				})
				return
			} else {
				c.Next()
			}
		}
	}
}