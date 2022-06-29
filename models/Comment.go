package models

import (
	"time"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	User_Id int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:User_Id"`
	Photo_Id int `json:"photo_id"`
	Photo Photo `json:"photo" gorm:"foreignKey:Photo_Id"`
	Message string `json:"message"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type CommentList struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	Photo_Id int `json:"photo_id"`
	User_Id int `json:"user_id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	User CommentListUser `json:"user"`
	Photo CommentListPhoto `json:"photo"`
}

type CommentListUser struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type CommentListPhoto struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	Photo_Url string `json:"photo_id"`
	User_Id int `json:"user_id"`
}