package routes

import (
	"grinder/pkg/logger"
	"grinder/pkg/repository"
	"net/http"

	"go.uber.org/zap"

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

func (ar *AppRouter) UpdatePerson(c *gin.Context) {
	userID, userVersion, valid := ar.getUserAndVersion(c)
	if !valid {
		return
	}

	var p struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}
	user, err := ar.userRepo.UpdateUser(repository.User{
		ID:        userID,
		Version:   userVersion,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
		return
	}
	user.Pass = ""
	user.Version = 0
	c.JSON(http.StatusOK, gin.H{"ok": true, "user": user})
}

func (ar *AppRouter) UpdatePersonPass(c *gin.Context) {
	userID, userVersion, valid := ar.getUserAndVersion(c)
	if !valid {
		return
	}

	var p struct {
		OldPass string `json:"old_pass"`
		NewPass string `json:"new_pass"`
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}
	user, valid, err := ar.userRepo.UpdateUserPassword(repository.UpdateUserPass{
		ID:      userID,
		Version: userVersion,
		Pass:    p.NewPass,
		OldPass: p.OldPass,
	})
	if err != nil {
		logger.Error("error change password", err, zap.String("old_pass", p.OldPass), zap.String("new_pass", p.NewPass), zap.Int64("user", userID))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
		return
	}

	if !valid {
		c.JSON(http.StatusOK, gin.H{"ok": false, "error": "wrong password"})
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
