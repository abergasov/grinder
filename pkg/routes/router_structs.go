package routes

import (
	"grinder/pkg/repository"
	"time"

	"github.com/gin-gonic/gin"
)

const AdminRights int64 = 1
const ModeratorRights int64 = 2

type UserUpdateRequest struct {
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
}

type ISessionManager interface {
	CreateSession(userID int64, version int64) (string, error)
	ValidateSession(string) (int64, int64)
	GetTokenLiveTime() time.Duration
	AuthMiddleware(*gin.Context)
	GetUserAndVersion(c *gin.Context) (int64, int64, bool)
}

type IUserRepo interface {
	RegisterUser(mail, password string) (registered int64, exist bool, err error)
	LoginUser(mail, password string) (userID int64, userVersion int64, err error)
	CheckVersion(userID, version int64) (valid bool, err error)
	GetUser(userID, version int64) (*repository.User, bool, error)
	UpdateUser(u *repository.User) error
	UpdateUserPassword(userID, userV int64, oldPass, newPass string) (*repository.User, bool, error)
}

type IPersonsRepo interface {
	LoadPersons(offset int64) ([]repository.Person, []repository.PersonRight, error)
	GetRightsMap() map[int64]string
	UpdateUser(userID int64, firstName, lastName, email string, active bool) (bool, error)
}

type IRightsChecker interface {
	CheckRight(rights []int64, ver func(*gin.Context) (int64, int64, bool)) gin.HandlerFunc
}
