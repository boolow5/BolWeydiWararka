package main

import (
	"os"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/boolow5/BolWeydi/controllers"
	"github.com/boolow5/BolWeydi/middlewares"
	_ "github.com/boolow5/BolWeydi/models"
)

func main() {
	router := gin.Default()
	// router.Use(middlewares.CheckDomain())
	jwtMiddleware := middlewares.NewJWTMiddleware()
	router.GET("/", controllers.Index)
	auth := router.Group("/")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome you're authorized to see this!"})
		})
	}
	router.POST("/user", controllers.AddUser)
	router.PUT("/user", controllers.UpdateUser)
	router.DELETE("/user", controllers.DeleteUser)
	//router.POST("/auth", controllers.AuthenticateUser)

	router.POST("/login", jwtMiddleware.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	router.Run(port)
}
