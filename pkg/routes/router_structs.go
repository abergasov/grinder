package routes

import (
	"grinder/pkg/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type ISessionManager interface {
	CreateSession(userID int64, version int64) (string, error)
	ValidateSession(string) (int64, int64)
	GetTokenLiveTime() time.Duration
	AuthMiddleware(*gin.Context)
	GetUserAndVersion(*gin.Context) (int64, int64, bool)
}

type IUserRepo interface {
	RegisterUser(mail, password string) (registered int64, exist bool, err error)
	LoginUser(mail, password string) (userID int64, userVersion int64, err error)
	CheckVersion(userID, version int64) (valid bool, err error)
	GetUser(userID, version int64) (*repository.User, bool, error)
	UpdateUser(u repository.User) (*repository.User, error)
	UpdateUserPassword(userID, userV int64, oldPass, newPass string) (*repository.User, bool, error)
}
