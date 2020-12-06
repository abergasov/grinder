package repository

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	jwtCookie = "test"
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

func TestSessionManager_GetUserAndVersion(t *testing.T) {
	s := InitSessionManager("abc", jwtCookie, 100*time.Millisecond)

	type cDT struct {
		uV           interface{}
		uID          interface{}
		expectedBool bool
		expectedUID  int64
		expectedUV   int64
	}
	cases := []cDT{
		{"adwdaw", "adwdaw", false, 0, 0},
		{nil, "adwdaw", false, 0, 0},
		{"adwdaw", nil, false, 0, 0},
		{int64(2), int64(1), true, int64(1), int64(2)},
		{int32(2), int64(1), false, 0, 0},
		{nil, int64(1), false, 0, 0},
		{int32(2), int32(1), false, 0, 0},
	}
	for _, cs := range cases {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		if cs.uID != nil {
			c.Set("user_id", cs.uID)
		}
		if cs.uV != nil {
			c.Set("user_version", cs.uV)
		}
		uID, uV, ok := s.GetUserAndVersion(c)
		if ok != cs.expectedBool {
			t.Errorf("expected %t, got %t here", cs.expectedBool, ok)
		}
		if uID != cs.expectedUID {
			t.Errorf("unexpected userID here, want %d, got %d", cs.expectedUID, uID)
		}
		if uV != cs.expectedUV {
			t.Errorf("unexpected userVersion here, want %d, got %d", cs.expectedUV, uV)
		}
	}

}
