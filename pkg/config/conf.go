package config

import (
	"grinder/pkg/logger"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ProdEnv bool   `yaml:"prod_env"`
	HostURL string `yaml:"host_url"`
	AppPort string `yaml:"app_port"`
	DBConf  DBConf `yaml:"db_conf"`
}

type DBConf struct {
	DBHost string `yaml:"db_host"`
	DBName string `yaml:"db_name"`
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
	DBPort string `yaml:"db_port"`
}

func InitConf(confFilePath string) *AppConfig {
	path, err := os.Getwd()
	if err != nil {
		logger.Fatal("Can't locate current dir", err)
	}

	confFile := path + confFilePath
	confFile = filepath.Clean(confFile)
	logger.Info("Try read config file", zap.String("path", confFile))

	file, errP := os.Open(confFile)
	if errP != nil {
		logger.Fatal("Can't open config file", errP)
	}
	defer file.Close()
	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.Fatal("Invalid config file", err)
	}

	return &cfg
}
