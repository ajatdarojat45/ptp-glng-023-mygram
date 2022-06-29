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
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}