package middlewares

import (
	"time"

	jwt "gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/gin-gonic/gin.v1"
)

func NewJWTMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: MyAuthenticator,
		Authorizator:  MyAuthorizor,
		Unauthorized:  GetOUT,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
	}
}

func MyAuthenticator(userId string, password string, c *gin.Context) (string, bool) {
	if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		return userId, true
	}

	return userId, false
}
func MyAuthorizor(userId string, c *gin.Context) bool {
	if userId == "admin" {
		return true
	}

	return false
}

func GetOUT(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
