package models

import (
	"time"

	"to-do-list-go/src/config"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `form:"name" json:"name" binding:"required"`
	Username  string    `form:"username" json:"username" binding:"required"`
	Email     string    `form:"email" json:"email" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

// Get all users
func GetAllUsers(user *[]User) (err error) {
	err = config.DB.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Get single user
func GetUser(user *User, id int) (err error) {
	err = config.DB.First(user, id).Error
	if err != nil {
		return err
	}
	return nil
}

// Create user
func CreateUser(user *User) (err error) {
	err = config.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
