package repository

import (
	"grinder/pkg/config"
	"grinder/pkg/storage"
)

type PersonsRepo struct {
	db *storage.DBConnector
}

func InitPersonsRepository(cnf *config.AppConfig, db *storage.DBConnector) *PersonsRepo {
	return &PersonsRepo{
		db: db,
	}
}

func (p *PersonsRepo) LoadPersons() {

}
