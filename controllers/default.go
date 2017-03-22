package controllers

import (
	"strconv"

	"github.com/boolow5/BolWeydi/models"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {

}

func Index(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Welcome to iWeydi"})
}

// Auth Controllers
func AddUser(context *gin.Context) {
	user := models.User{}

	context.BindJSON(&user)
	err := user.SetPassword(user.Password)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if user.Profile == nil {
		user.Profile = &models.Profile{FirstName: user.Username}
	}
	var saved bool
	saved, err = user.Add()
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	if !saved {
		context.JSON(200, gin.H{"warning": "failed without warning"})
	}

	context.JSON(200, gin.H{"success": "Saved successfully"})
}

func UpdateUser(context *gin.Context) {
	id, got := context.GetQuery("user_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing user id"})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if userId < 1 {
		context.JSON(200, gin.H{"error": "invalid user id"})
		return
	}
	_, user := models.GetUserById(userId)
	err = context.BindJSON(&user)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	saved, err := user.Update()
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if !saved {
		context.JSON(200, gin.H{"warning": "update failed without error"})
		return
	}
	context.JSON(200, gin.H{"success": "updated successfully"})
}

func DeleteUser(context *gin.Context) {
	id, got := context.GetQuery("user_id")
	if !got {
		context.JSON(200, gin.H{"error": "missing user id"})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	if userId < 1 {
		context.JSON(200, gin.H{"error": "invalid user id"})
		return
	}
	_, user := models.GetUserById(userId)
	deleted, err := user.Delete()
	data := map[string]interface{}{}
	if err != nil {
		data["error"] = err.Error()
		context.JSON(200, data)
		return
	}
	if !deleted {
		data["message"] = "User not deleted"
		context.JSON(200, data)
		return
	}
	data["message"] = "User deleted successfully"
	context.JSON(200, data)
}

func AuthenticateUser(context *gin.Context) {
	user := models.User{Username: "boolow5", Password: "sharaf.143"}
	err, loggedinUser := user.Authenticate()
	error_message := ""
	if err != nil {
		error_message = err.Error()
	}
	context.JSON(200, gin.H{"user": loggedinUser, "error": error_message})
}
