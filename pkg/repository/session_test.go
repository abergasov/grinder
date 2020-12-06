package repository

import (
	"fmt"
	"net/http"
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

func TestSessionManager_AuthMiddleware(t *testing.T) {
	s := InitSessionManager("abc", jwtCookie, 222*time.Second)
	token, err := s.CreateSession(666, 666)
	if err != nil {
		t.Error("unexpected error here", err.Error())
	}

	type tM struct {
		cookieName   string
		cookieVal    string
		expectedCode int
		expectedUid  bool
		expectedUV   bool
	}

	cases := []tM{
		{"asd", "asds", 401, false, false},
		{jwtCookie, "asds", 401, false, false},
		{jwtCookie, token, 200, true, true},
	}

	for _, cs := range cases {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.Header.Set("Cookie", fmt.Sprintf("%s=%s", cs.cookieName, cs.cookieVal))

		s.AuthMiddleware(c)
		if c.Writer.Status() != cs.expectedCode {
			t.Errorf("expected %d, got %d", cs.expectedCode, c.Writer.Status())
		}
		if cs.expectedUV {
			v, exist := c.Get("user_version")
			if !exist {
				t.Error("expected user version here")
			}
			if v.(int64) != 666 {
				t.Errorf("invalid user version, vant %d, got %d", 666, v)
			}
		}

		if cs.expectedUid {
			v, exist := c.Get("user_id")
			if !exist {
				t.Error("expected user id here")
			}
			if v.(int64) != 666 {
				t.Errorf("invalid user id, vant %d, got %d", 666, v)
			}
		}
	}

}
