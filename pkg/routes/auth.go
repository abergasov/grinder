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

var (
	reEmail   = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	jwtCookie = "rc"
)

func (ar *AppRouter) LoginUser(c *gin.Context) {
	u := ar.checkAuthRequest(c)
	if u == nil {
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

	ar.setSecretCookie(c, jwtCookie, token)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}

func (ar *AppRouter) RegisterUser(c *gin.Context) {
	u := ar.checkAuthRequest(c)
	if u == nil {
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

	ar.setSecretCookie(c, jwtCookie, token)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}

func (ar *AppRouter) RefreshToken(c *gin.Context) {
	token, err := c.Cookie(jwtCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
		return
	}
	userID, userVersion := ar.sessionRepo.ValidateSession(token)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
		return
	}

	valid, err := ar.userRepo.CheckVersion(userID, userVersion)
	if err != nil {
		logger.Error("error while check token", err, zap.Int64("userId", userID), zap.String("token", token))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while refresh"})
		return
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
		return
	}

	token, err = ar.sessionRepo.CreateSession(userID, userVersion)
	if err != nil {
		logger.Error("error while create token", err, zap.Int64("userId", userID), zap.String("token", token))
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "error while login"})
		return
	}

	ar.setSecretCookie(c, jwtCookie, token)
	c.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}

func (ar *AppRouter) checkAuthRequest(c *gin.Context) *regioRequesto {
	var u regioRequesto
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return nil
	}

	if !reEmail.MatchString(u.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid email"})
		return nil
	}
	return &u
}

func (ar *AppRouter) Logout(c *gin.Context) {
	ar.setSecretCookie(c, jwtCookie, "")
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (ar *AppRouter) setSecretCookie(c *gin.Context, keyName, keyValue string) {
	liveTime := repository.GetTokenLiveTime()
	path := "/api/"
	c.SetCookie(keyName, keyValue, int(liveTime), path, ar.config.HostURL, ar.config.SSLEnable, true)
}
