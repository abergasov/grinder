package routes

import (
	"errors"
	"grinder/pkg/repository"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAppRouter_GetPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	engine, token := createRouter(ctrl)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(0), int64(0), false)
	uRepo.EXPECT().GetUser(tUser, tUserV)
	w := performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusBadRequest, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(tUser, tUserV, true)
	uRepo.EXPECT().GetUser(tUser, tUserV).Return(nil, false, errors.New("42"))
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusUnauthorized, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(tUser, tUserV, true)
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusInternalServerError, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(tUser, tUserV, true)
	uRepo.EXPECT().GetUser(tUser, tUserV).Return(&repository.User{}, true, nil)
	w = performRequest(engine, "POST", "/api/data/profile", token, nil)
	compareCode(http.StatusOK, w, t)
}

func TestAppRouter_UpdatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	engine, token := createRouter(ctrl)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(int64(0), int64(0), false)
	w := performRequest(engine, "POST", "/api/data/profile/update", token, nil)
	compareCode(http.StatusBadRequest, w, t)

	request := &updateRequest{
		NewPass: "1",
		OldPass: "2",
	}

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(tUser, tUserV, true)
	uRepo.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("42"))
	w = performRequest(engine, "POST", "/api/data/profile/update", token, request)
	compareCode(http.StatusInternalServerError, w, t)

	session.EXPECT().AuthMiddleware(gomock.Any())
	session.EXPECT().GetUserAndVersion(gomock.Any()).Return(tUser, tUserV, true)
	uRepo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	w = performRequest(engine, "POST", "/api/data/profile/update", token, request)
	compareCode(http.StatusOK, w, t)
}
