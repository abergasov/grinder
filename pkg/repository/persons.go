package repository

import (
	"grinder/pkg/config"
	"grinder/pkg/storage"
)

type PersonsRepo struct {
	db *storage.DBConnector
}

type person struct {
	UserID    int64  `db:"user_id" json:"user_id"`
	Email     string `db:"email" json:"email"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func InitPersonsRepository(cnf *config.AppConfig, db *storage.DBConnector) *PersonsRepo {
	return &PersonsRepo{
		db: db,
	}
}

func (p *PersonsRepo) LoadPersons() {

}
