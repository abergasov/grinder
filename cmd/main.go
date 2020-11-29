package main

import (
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"grinder/pkg/routes"
	"grinder/pkg/storage"

	"go.uber.org/zap"
)

var (
	appName   = "el_grinder"
	buildTime = "_dev"
	buildHash = "_dev"
)

func main() {
	logger.NewLogger()
	appConf := config.InitConf("/configs/conf.yaml")
	dbConnect := storage.InitDBConnect(appConf)
	router := routes.InitRouter(appConf, dbConnect, appName, buildTime, buildHash)

	logger.Info("Server running on port", zap.String("port", appConf.AppPort), zap.String("url", "http://localhost:"+appConf.AppPort))
	err := router.InitRoutes().Run(":" + appConf.AppPort)
	if err != nil {
		logger.Fatal("Router error", err)
	}
}
