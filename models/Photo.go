package models

import (
	"time"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:User_Id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}