package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"grinder/pkg/repository"
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

type updateRequest struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

func performRequest(r http.Handler, method, path, token string, payload interface{}) *httptest.ResponseRecorder {
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

//go:generate mockgen -source=router_structs.go -destination=auth_mock.go -package=routes
func TestAppRouter_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf := getSampleConf()
	//session := repository.InitSessionManager("abc", jwtCookieExp)
	session := NewMockISessionManager(ctrl)
	uRepo := NewMockIUserRepo(ctrl)
	router := InitRouter(conf, &RouterConfig{
		UserRepo:    uRepo,
		SessionRepo: session,
	}, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()
	w := performRequest(engine, "POST", "/api/auth/login", "", nil)
	compareCode(http.StatusBadRequest, w, t)

	body := &requesto{
		Email:      "dawdawdwa",
		Password:   "adwadawd",
		RePassword: "1",
	}
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusBadRequest, w, t)

	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusBadRequest, w, t)

	// valid data here
	body.Email = "a@a.aa"
	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", nil)
	session.EXPECT().GetTokenLiveTime()
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusOK, w, t)

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusInternalServerError, w, t)

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(666), int64(666), errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusInternalServerError, w, t)

	uRepo.EXPECT().LoginUser(body.Email, body.Password).Return(int64(0), int64(0), nil)
	w = performRequest(engine, "POST", "/api/auth/login", "", body)
	compareCode(http.StatusUnauthorized, w, t)

}

func TestAppRouter_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf := getSampleConf()
	//session := repository.InitSessionManager("abc", jwtCookieExp)
	session := NewMockISessionManager(ctrl)
	uRepo := NewMockIUserRepo(ctrl)
	router := InitRouter(conf, &RouterConfig{
		UserRepo:    uRepo,
		SessionRepo: session,
	}, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()

	w := performRequest(engine, "POST", "/api/auth/register", "", nil)
	compareCode(http.StatusBadRequest, w, t)

	body := &requesto{
		Email:      "dawdawdwa",
		Password:   "adwadawd",
		RePassword: "1",
	}

	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusBadRequest, w, t)

	body.Email = "a@a.au"

	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusBadRequest, w, t)

	body.RePassword = "adwadawd"
	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(0), false, errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusInternalServerError, w, t)

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(0), true, nil)
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusConflict, w, t)

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(666), false, nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", nil)
	session.EXPECT().GetTokenLiveTime()
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusOK, w, t)

	uRepo.EXPECT().RegisterUser(body.Email, body.Password).Return(int64(666), false, nil)
	session.EXPECT().
		CreateSession(gomock.Any(), gomock.Any()).
		Return("abc", errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/register", "", body)
	compareCode(http.StatusInternalServerError, w, t)
}

func TestAppRouter_RefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf := getSampleConf()
	sessionReal := repository.InitSessionManager("abc", jwtCookie, jwtCookieExp)
	token, err := sessionReal.CreateSession(666, 666)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}
	session := NewMockISessionManager(ctrl)
	uRepo := NewMockIUserRepo(ctrl)
	router := InitRouter(conf, &RouterConfig{
		UserRepo:    uRepo,
		SessionRepo: session,
	}, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()

	w := performRequest(engine, "POST", "/api/auth/refresh", "", nil)
	compareCode(http.StatusUnauthorized, w, t)

	session.EXPECT().ValidateSession(gomock.Any()).Return(int64(0), int64(0))
	w = performRequest(engine, "POST", "/api/auth/refresh", token, nil)
	compareCode(http.StatusUnauthorized, w, t)

	session.EXPECT().ValidateSession(gomock.Any()).Return(int64(666), int64(666))
	uRepo.EXPECT().CheckVersion(gomock.Any(), gomock.Any()).Return(false, errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/refresh", token, nil)
	compareCode(http.StatusInternalServerError, w, t)

	session.EXPECT().ValidateSession(gomock.Any()).Return(int64(666), int64(666))
	uRepo.EXPECT().CheckVersion(gomock.Any(), gomock.Any()).Return(false, nil)
	w = performRequest(engine, "POST", "/api/auth/refresh", token, nil)
	compareCode(http.StatusUnauthorized, w, t)

	session.EXPECT().ValidateSession(gomock.Any()).Return(int64(666), int64(666))
	uRepo.EXPECT().CheckVersion(gomock.Any(), gomock.Any()).Return(true, nil)
	session.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Return("", errors.New("42"))
	w = performRequest(engine, "POST", "/api/auth/refresh", token, nil)
	compareCode(http.StatusInternalServerError, w, t)

	session.EXPECT().ValidateSession(gomock.Any()).Return(int64(666), int64(666))
	uRepo.EXPECT().CheckVersion(gomock.Any(), gomock.Any()).Return(true, nil)
	session.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Return("avc", nil)
	session.EXPECT().GetTokenLiveTime()
	w = performRequest(engine, "POST", "/api/auth/refresh", token, nil)
	compareCode(http.StatusOK, w, t)
}

func compareCode(expected int, w *httptest.ResponseRecorder, t *testing.T) {
	if w.Code != expected {
		t.Errorf("expected %d code, got %d", expected, w.Code)
	}
}
