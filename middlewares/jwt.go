package middlewares

import (
	"time"

	"github.com/boolow5/BolWeydi/models"
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

func MyAuthenticator(username string, password string, c *gin.Context) (string, bool) {
	user := &models.User{Username: username, Password: password}
	err, user := user.Authenticate()
	if err != nil {
		return "", false
	}

	return user.Username, true
}
func MyAuthorizor(userId string, c *gin.Context) bool {
	user := models.User{Username: userId}
	allow, err := user.Authorize("*")
	if err != nil {
		return false
	}
	if !allow {
		return false
	}

	return true
}

func GetOUT(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
