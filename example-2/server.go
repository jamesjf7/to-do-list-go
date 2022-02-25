package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

var users = []User{}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func registerUserHandler(c *gin.Context) {
	// best practice to get the body of the request
	user := User{
		ID:       len(users) + 1,
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
	}

	if err := c.ShouldBindWith(&user, binding.Form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
			"error":   err.Error(),
		})
		return
	}

	users = append(users, user)
	fmt.Println(user)

	c.JSON(200, gin.H{
		"message": "user created",
		"user":    &user,
	})
}

func getAllUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
	fmt.Println(users)
}

func getUserHandler(c *gin.Context) {
	username := c.Param("username")

	var user_found = User{}
	// get user with username
	for _, user := range users {
		if user.Username == username {
			user_found = user
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user_found,
	})
}

func updateUserHandler(c *gin.Context) {
	username := c.Param("username")

	var user_found = User{}
	// get user with username
	for _, user := range users {
		if user.Username == username {
			user_found = user
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user":    user_found,
	})
}

func main() {
	app := gin.Default() // app in express

	// routes
	app.GET("/", indexHandler)
	app.POST("/users", registerUserHandler)
	app.GET("/users/", getAllUsersHandler)
	app.GET("/users/:username", getUserHandler)
	app.PUT("/users/:username", updateUserHandler)

	app.Run(":8080")
}
