package controllers

import (
	"fmt"

	"github.com/boolow5/BolWeydi/models"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	fmt.Println("initializing controllers...")
}

func Index(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Welcome to iWeydi"})
}

// Auth Controllers
func AddUser(context *gin.Context) {
	user := models.User{}
	var err_message string
	var saved bool
	user.Username = "boolow5"
	err := user.SetPassword("sharaf.143")
	if err == nil {
		user.Role = "admin"
		user.Profile = &models.Profile{FirstName: "Mahdi", MiddleName: "Ahmed", LastName: "Bolow", AnswerCount: 100, AnswerViewCount: 3545}
		saved, err = user.Add()
	}
	if err != nil {
		err_message = err.Error()
	}

	context.JSON(200, gin.H{"user": user, "saved": saved, "error": err_message})
}

func UpdateUser(context *gin.Context) {
	_, user := models.GetUserById(1)
	user.Username = "mahdiyare"
	saved, err := user.Update()
	var err_message string
	if err != nil {
		err_message = err.Error()
	}
	context.JSON(200, gin.H{"saved": saved, "error": err_message})
}

func DeleteUser(context *gin.Context) {
	_, user := models.GetUserById(1)
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
