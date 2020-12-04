package repository

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

type SessionManager struct {
	jwtKey           []byte
	jwtCookie        string
	sessionValidTime time.Duration
}

type sessionClaims struct {
	UserID      int64 `json:"user_id"`
	UserVersion int64 `json:"v"`
	jwt.StandardClaims
}

func InitSessionManager(jwtKey, jwtCookie string, tokenLiveTime time.Duration) *SessionManager {
	return &SessionManager{
		jwtKey:           []byte(jwtKey),
		jwtCookie:        jwtCookie,
		sessionValidTime: tokenLiveTime,
	}
}

func (s *SessionManager) GetTokenLiveTime() time.Duration {
	return s.sessionValidTime
}

func (s *SessionManager) CreateSession(userID, userVersion int64) (string, error) {
	atClaims := sessionClaims{
		UserID:      userID,
		UserVersion: userVersion,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.sessionValidTime).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	return at.SignedString(s.jwtKey)
}

func (s *SessionManager) ValidateSession(sessionId string) (userID, version int64) {
	token, err := jwt.ParseWithClaims(sessionId, &sessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return
	}

	if !token.Valid {
		return
	}

	if claims, ok := token.Claims.(*sessionClaims); ok {
		return claims.UserID, claims.UserVersion
	}

	return
}

func (s *SessionManager) GetUserAndVersion(c *gin.Context) (int64, int64, bool) {
	val, exist := c.Get("user_id")
	if !exist {
		return 0, 0, false
	}
	userID, ok := val.(int64)
	if !ok {
		return 0, 0, false
	}
	val, exist = c.Get("user_version")
	if !exist {
		return 0, 0, false
	}
	userVersion, ok := val.(int64)
	if !ok {
		return 0, 0, false
	}
	return userID, userVersion, true
}

func (s *SessionManager) AuthMiddleware(c *gin.Context) {
	token, err := c.Cookie(s.jwtCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
		return
	}
	userID, userVersion := s.ValidateSession(token)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "expired"})
		return
	}

	c.Set("user_id", userID)
	c.Set("user_version", userVersion)
	c.Next()
}
