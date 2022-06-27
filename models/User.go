package models
import (
	"gorm.io/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
	"fmt"	
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		fmt.Println("Failed to encrypt password: ", err)
		return err
	}

	u.Password = string(pwd)
	return nil
}