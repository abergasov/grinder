package middleware

import (
	"grinder/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtCookie = ""
var session *repository.SessionManager

func InitMiddleware(jCookie string, sM *repository.SessionManager) {
	jwtCookie = jCookie
	session = sM
}

func AuthOrchestraMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(jwtCookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
			return
		}
		userID, userVersion := session.ValidateSession(token)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
			return
		}

		c.Set("user_id", userID)
		c.Set("user_version", userVersion)
		c.Next()
	}
}
