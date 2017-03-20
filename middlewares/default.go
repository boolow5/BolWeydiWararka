package middlewares

import (
	"errors"
	"log"

	"github.com/boolow5/BolWeydi/g"
	"gopkg.in/gin-gonic/gin.v1"
)

func CheckDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Print("Checking ip = from: \"" + c.Request.Host + "\"...")
		if !AuthorizedIP(c.Request.Host) {
			c.JSON(403, gin.H{"error": "Access Denied"})
			c.AbortWithError(403, errors.New("You're from unauthorized application from:"+c.Request.Host))
			return
		}
		log.Print("OK\n")
		c.Next()
	}
}

func AuthorizedIP(ip string) bool {
	for i := 0; i < len(g.KNOWN_IP_ADDRESSES); i++ {
		if ip == g.KNOWN_IP_ADDRESSES[i] {
			return true
		}
	}
	return false
}
