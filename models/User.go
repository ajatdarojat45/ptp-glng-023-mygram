package models
import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}