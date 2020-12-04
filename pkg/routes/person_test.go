package routes

import (
	"errors"
	"grinder/pkg/logger"
	"grinder/pkg/repository"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAppRouter_GetPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger.NewLogger()
	defer ctrl.Finish()
	conf := getSampleConf()
	sessionReal := repository.InitSessionManager("abc", jwtCookie, jwtCookieExp)
	session := NewMockISessionManager(ctrl)
	token, err := sessionReal.CreateSession(666, 666)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	uRepo := NewMockIUserRepo(ctrl)
	session.EXPECT().AuthMiddleware(gomock.Any())
	router := InitRouter(conf, uRepo, session, jwtCookie, "test", "test", "hash")
	engine := router.InitRoutes()

	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(0), int64(0), false)
	uRepo.EXPECT().GetUser(int64(666), int64(666))
	w := performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusBadRequest, w, t)

	if err != nil {
		t.Error("unexpected error", err.Error())
	}
	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(666), int64(666), true)
	uRepo.EXPECT().GetUser(int64(666), int64(666)).Return(nil, false, errors.New("42"))
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusUnauthorized, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(666), int64(666), true)
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusInternalServerError, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(666), int64(666), true)
	uRepo.EXPECT().GetUser(int64(666), int64(666)).Return(&repository.User{}, true, nil)
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusOK, w, t)
}
