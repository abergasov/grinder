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
	sessionRepo   ISessionManager
	rightsChecker IRightsChecker
}

type RouterConfig struct {
	UserRepo    IUserRepo
	RightsRepo  IRightsChecker
	SessionRepo ISessionManager
}

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
		sessionRepo:   rCnf.SessionRepo,
		rightsChecker: rCnf.RightsRepo,
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

	authorizedDataGroup := ar.GinEngine.Group("/api/data")
	authorizedDataGroup.Use(ar.sessionRepo.AuthMiddleware)
	authorizedDataGroup.POST("profile", ar.GetPerson)
	authorizedDataGroup.POST("profile/update", ar.UpdatePerson)
	authorizedDataGroup.POST("profile/update_password", ar.UpdatePersonPass)
	authorizedDataGroup.
		Use(ar.rightsChecker.CheckRight([]int64{AdminRights}, ar.sessionRepo.GetUserAndVersion)).
		POST("users/list", ar.UpdatePersonPass)
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
