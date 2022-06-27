package config

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"fmt"
	"mygram/models"
)

func DBConnect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to cennect to database")
	}
	fmt.Println("Success connect to DB using GORM")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.SocialMedia{})
	return db
}