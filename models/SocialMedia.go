package models

import (
	"gorm.io/gorm"
	"time"
)

type SocialMedia struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:User_Id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type SocialMediaList struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id int `json:"user_id"`
	User SocialMediaListUser `json:"user"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type SocialMediaListUser struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

