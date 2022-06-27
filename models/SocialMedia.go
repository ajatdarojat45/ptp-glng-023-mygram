package models

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	UserId int `json:"user_id"`
}