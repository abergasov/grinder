package main

import (
	"grinder/pkg/config"
	"grinder/pkg/logger"
)

func main() {
	logger.NewLogger()
	appConf := config.InitConf("/configs/conf.yaml")
	println(appConf.AppPort)
}
