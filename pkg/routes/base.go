package routes

import (
	"grinder/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	GinEngine     *gin.Engine
	config        *config.AppConfig
	AppName       string
	AppBuildTime  string
	AppBuildHash  string
	jwtCookie     string
	userRepo      IUserRepo
	personsRepo   IPersonsRepo
	sessionRepo   ISessionManager
	rightsChecker IRightsChecker
}

type RouterConfig struct {
	UserRepo    IUserRepo
	RightsRepo  IRightsChecker
	PersonsRepo IPersonsRepo
	SessionRepo ISessionManager
}

var adminPagesRights = []int64{AdminRights}

func InitRouter(cnf *config.AppConfig, rCnf *RouterConfig, jwtCookie, appName, appBuild, appHash string) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	router := &AppRouter{
		GinEngine:     gin.Default(),
		config:        cnf,
		AppName:       appName,
		AppBuildHash:  appHash,
		AppBuildTime:  appBuild,
		jwtCookie:     jwtCookie,
		userRepo:      rCnf.UserRepo,
		rightsChecker: rCnf.RightsRepo,
		personsRepo:   rCnf.PersonsRepo,
		sessionRepo:   rCnf.SessionRepo,
	}
	return router
}

func (ar *AppRouter) InitRoutes() *gin.Engine {
	ar.GinEngine.GET("/ping", ar.Ping)
	authGroup := ar.GinEngine.Group("/api/auth")
	authGroup.POST("login", ar.LoginUser)
	authGroup.POST("register", ar.RegisterUser)
	authGroup.POST("refresh", ar.RefreshToken)
	authGroup.POST("logout", ar.Logout)

	authDataGroup := ar.GinEngine.Group("/api/data")
	authDataGroup.Use(ar.sessionRepo.AuthMiddleware)
	authDataGroup.POST("profile", ar.GetPerson)
	authDataGroup.POST("profile/update", ar.UpdatePerson)
	authDataGroup.POST("profile/update_password", ar.UpdatePersonPass)
	authDataGroup.
		Use(ar.rightsChecker.CheckRight(adminPagesRights, ar.sessionRepo.GetUserAndVersion)).
		POST("users/list", ar.GetUsersList)
	authDataGroup.
		Use(ar.rightsChecker.CheckRight(adminPagesRights, ar.sessionRepo.GetUserAndVersion)).
		POST("users/roles/list", ar.GetUsersRoles)
	authDataGroup.
		Use(ar.rightsChecker.CheckRight(adminPagesRights, ar.sessionRepo.GetUserAndVersion)).
		POST("users/update", ar.UpdateUser)
	return ar.GinEngine
}

func (ar *AppRouter) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":         true,
		"build_hash": ar.AppBuildHash,
		"build_time": ar.AppBuildTime,
		"app_name":   ar.AppName,
	})
}
