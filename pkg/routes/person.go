package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *AppRouter) GetPerson(c *gin.Context) {
	userID, userVersion, valid := ar.getUserAndVersion(c)
	if !valid {
		return
	}

	user, versionCorrect, err := ar.userRepo.GetUser(userID, userVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
		return
	}
	if !versionCorrect {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid key"})
		return
	}
	user.Pass = ""
	user.Version = 0
	c.JSON(http.StatusOK, gin.H{"ok": true, "user": user})
}

func (ar *AppRouter) getUserAndVersion(c *gin.Context) (int64, int64, bool) {
	val, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid user"})
		return 0, 0, false
	}
	userID, ok := val.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid user"})
		return 0, 0, false
	}
	val, exist = c.Get("user_version")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid user"})
		return 0, 0, false
	}
	userVersion, ok := val.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid user"})
		return 0, 0, false
	}
	return userID, userVersion, true
}
