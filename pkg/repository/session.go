package repository

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SessionManager struct {
	jwtKey           []byte
	sessionValidTime time.Duration
}

type sessionClaims struct {
	UserID      int64 `json:"user_id"`
	UserVersion int64 `json:"v"`
	jwt.StandardClaims
}

func InitSessionManager(jwtKey string, tokenLiveTime time.Duration) *SessionManager {
	return &SessionManager{
		jwtKey:           []byte(jwtKey),
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
