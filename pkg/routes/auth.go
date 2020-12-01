package routes

import (
	"grinder/pkg/logger"
	"grinder/pkg/repository"
	"net/http"
	"regexp"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type regioRequesto struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (ar *AppRouter) LoginUser(c *gin.Context) {
	var u regioRequesto
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}
	if !reEmail.MatchString(u.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid email"})
		return
	}

	userID, version, err := ar.userRepo.LoginUser(u.Email, u.Password)
	if err != nil {
		logger.Error("error while login", err, zap.String("email", u.Email), zap.String("pass", u.Password))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while login"})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "wrong login/pass"})
		return
	}
	token, err := ar.sessionRepo.CreateSession(userID, version)
	if err != nil {
		logger.Error("error while create token", err, zap.String("email", u.Email), zap.String("pass", u.Password))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}

func (ar *AppRouter) RegisterUser(c *gin.Context) {
	var u regioRequesto
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}
	if u.RePassword != u.Password {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "password mismatch"})
		return
	}

	if !reEmail.MatchString(u.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid email"})
		return
	}

	userID, exist, err := ar.userRepo.RegisterUser(u.Email, u.Password)
	if err != nil {
		logger.Error("error while register", err, zap.String("email", u.Email), zap.String("pass", u.Password))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while register"})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{"ok": false, "error": "user already exist"})
		return
	}

	token, err := ar.sessionRepo.CreateSession(userID, repository.DefaultUserVersion)
	if err != nil {
		logger.Error("error while create token", err, zap.String("email", u.Email), zap.String("pass", u.Password))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}
