package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"mygram/config"
	"mygram/helpers"
	"mygram/models"
	"strings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("Bad Authorization"))
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("Bad Authorization"))
			return
		}
		if authSplit[0] != "Bearer" {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("Bad Authorization"))
			return
		}

		token, err := helpers.VerifyToken(authSplit[1])
		if err != nil {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("Bad Authorization"))
			return
		}

		dbResult := models.User{}
		username := token["username"]
		err = config.DBConnect().Debug().Where("username = ?", username).Last(&dbResult).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Username not found")
				c.AbortWithError(http.StatusBadRequest, errors.New("Username not found"))
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Set("userId", strconv.FormatInt(int64(dbResult.ID), 10))
		c.Next()
	}
}
