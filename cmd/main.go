package main

import (
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"grinder/pkg/repository"
	"grinder/pkg/routes"
	"grinder/pkg/storage"
	"time"

	"go.uber.org/zap"
)

var (
	appName   = "el_grinder"
	buildTime = "_dev"
	buildHash = "_dev"
	jwtCookie = "rc"
)

func main() {
	logger.NewLogger()
	appConf := config.InitConf("/configs/conf.yaml")
	dbConnect := storage.InitDBConnect(appConf)

	routerConf := &routes.RouterConfig{
		SessionRepo: repository.InitSessionManager(appConf.JWTKey, jwtCookie, 200*time.Minute),
		UserRepo:    repository.InitUserRepository(appConf, dbConnect),
		RightsRepo:  repository.InitRightManager(appConf, dbConnect),
		PersonsRepo: repository.InitPersonsRepository(appConf, dbConnect),
	}
	router := routes.InitRouter(appConf, routerConf, jwtCookie, appName, buildTime, buildHash)

	logger.Info("Server running on port", zap.String("port", appConf.AppPort), zap.String("url", "http://localhost:"+appConf.AppPort))
	err := router.InitRoutes().Run(":" + appConf.AppPort)
	if err != nil {
		logger.Fatal("Router error", err)
	}
}
