package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var (
	jwtCookie    = "test"
	jwtCookieExp = 2 * time.Minute
)

type requesto struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func performRequest(r http.Handler, method, path, token string, payload *requesto) *httptest.ResponseRecorder {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	if token != "" {
		cookie := http.Cookie{Name: jwtCookie, Value: token, Expires: time.Now().Add(jwtCookieExp)}
		req.AddCookie(&cookie)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func getSampleConf() *config.AppConfig {
	logger.NewLogger()
	dbConf := config.DBConf{
		DBPort: "",
		DBHost: "",
		DBPass: "",
		DBUser: "",
	}
	return &config.AppConfig{
		ProdEnv:   true,
		AppPort:   "8080",
		HostURL:   "",
		SSLEnable: false,
		JWTKey:    "",
		DBConf:    dbConf,
	}
}

//go:generate mockgen -source=router_structs.go -destination=auth_mock.go -package=routes IUserRepo
func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf := getSampleConf()
	//session := repository.InitSessionManager("abc", jwtCookieExp)
	session := NewMockISessionManager(ctrl)
	uRepo := NewMockIUserRepo(ctrl)
	router := InitRouter(conf, uRepo, session, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()
	w := performRequest(engine, "POST", "/api/auth/login", "", nil)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	body := &requesto{
		Email:      "dawdawdwa",
		Password:   "adwadawd",
		RePassword: "1",
	}
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	// valid data here
	body.Email = "a@a.aa"
	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", nil)
	session.EXPECT().GetTokenLiveTime()
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 code, got %d", w.Code)
	}

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 code, got %d", w.Code)
	}

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 code, got %d", w.Code)
	}

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(0), int64(0), nil)
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401 code, got %d", w.Code)
	}

}

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf := getSampleConf()
	//session := repository.InitSessionManager("abc", jwtCookieExp)
	session := NewMockISessionManager(ctrl)
	uRepo := NewMockIUserRepo(ctrl)
	router := InitRouter(conf, uRepo, session, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()

	w := performRequest(engine, "POST", "/api/auth/register", "", nil)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	body := &requesto{
		Email:      "dawdawdwa",
		Password:   "adwadawd",
		RePassword: "1",
	}

	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	body.Email = "a@a.au"

	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 code, got %d", w.Code)
	}

	body.RePassword = "adwadawd"
	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(0), false, errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 code, got %d", w.Code)
	}

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(0), true, nil)
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusConflict {
		t.Errorf("expected 409 code, got %d", w.Code)
	}

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(666), false, nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", nil)
	session.EXPECT().GetTokenLiveTime()
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 code, got %d", w.Code)
	}

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(666), false, nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 code, got %d", w.Code)
	}
}
