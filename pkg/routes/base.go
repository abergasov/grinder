package routes

import (
	"grinder/pkg/config"
	"grinder/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	GinEngine    *gin.Engine
	config       *config.AppConfig
	AppName      string
	AppBuildTime string
	AppBuildHash string
}

func InitRouter(cnf *config.AppConfig, dbConnect *storage.DBConnector, appName, appBuild, appHash string) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	router := &AppRouter{
		GinEngine:    gin.Default(),
		config:       cnf,
		AppName:      appName,
		AppBuildHash: appHash,
		AppBuildTime: appBuild,
	}
	return router
}

func (ar *AppRouter) InitRoutes() *gin.Engine {
	ar.GinEngine.GET("/ping", ar.Ping)
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
