package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primary_key"`
	UserID      int       `json:"user_id"`
	Name        string    `form:"name" json:"name" binding:"required"`
	Description string    `form:"description" json:"description" binding:"required"`
	IsDone      bool      `json:"is_done" json:"is_done"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" json:"updated_at"`
}
