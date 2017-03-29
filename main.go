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

		auth.PUT("/user", controllers.UpdateUser)
		auth.DELETE("/user", controllers.DeleteUser)
		//auth.POST("/auth", controllers.AuthenticateUser)

		auth.POST("/discussion", controllers.AddDiscussion)
		auth.PUT("/discussion", controllers.UpdateDiscussion)
		auth.DELETE("/discussion", controllers.DeleteDiscussion)

		auth.POST("/question", controllers.AddQuestion)
		auth.PUT("/question", controllers.UpdateQuestion)
		auth.DELETE("/question", controllers.DeleteQuestion)

		auth.POST("/answer", controllers.AddAnswer)
		auth.PUT("/answer", controllers.UpdateAnswer)
		auth.DELETE("/answer", controllers.DeleteAnswer)
	}
	router.POST("/user", controllers.AddUser)
	router.POST("/login", jwtMiddleware.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	router.Run(port)
}
