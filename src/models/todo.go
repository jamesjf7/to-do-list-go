package models

type Task struct {
	ID       int    `json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}
