package repository

import (
	"testing"
	"time"
)

var (
	jwtCookie    = "test"
	jwtCookieExp = 2 * time.Minute
)

func TestCreateSession(t *testing.T) {
	s := InitSessionManager("abc", jwtCookie, 100*time.Millisecond)

	token, err := s.CreateSession(666, 666)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	userID, version := s.ValidateSession(token)
	if userID != 666 {
		t.Errorf("mismatch userId, want 666, got %d", userID)
	}

	if version != 666 {
		t.Errorf("mismatch version, want 666, got %d", version)
	}

	time.Sleep(2 * time.Second)
	userID, version = s.ValidateSession(token)
	if userID != 0 || version != 0 {
		t.Errorf("token expired, but check return userId, user %d, version %d", userID, version)
	}
}
