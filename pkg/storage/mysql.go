package storage

import (
	"fmt"
	"grinder/pkg/config"
	"grinder/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConnector struct {
	Client *sqlx.DB
}

func InitDBConnect(cnf *config.AppConfig) *DBConnector {
	conStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", cnf.DBConf.DBUser, cnf.DBConf.DBPass, cnf.DBConf.DBHost, cnf.DBConf.DBPort, cnf.DBConf.DBName)
	db, err := sqlx.Connect("mysql", conStr)
	if err != nil {
		logger.Fatal("error connect to db", err)
	}
	return &DBConnector{
		Client: db,
	}
}
