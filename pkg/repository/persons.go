package repository

import (
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"grinder/pkg/storage"

	"github.com/jmoiron/sqlx/types"

	"github.com/jmoiron/sqlx"
)

type PersonsRepo struct {
	db           *storage.DBConnector
	systemRights map[int64]string
}

const loadLimit = 500

type Person struct {
	UserID     int64         `db:"user_id" json:"user_id"`
	Email      string        `db:"email" json:"email"`
	FirstName  string        `db:"first_name" json:"first_name"`
	LastName   string        `db:"last_name" json:"last_name"`
	RegisterAt string        `db:"register_date" json:"register_date"`
	Active     types.BitBool `db:"active,bit" json:"active"`
	Rights     []int64       `db:"right_id" json:"right_id"`
}

type PersonRight struct {
	UserID  int64 `db:"user_id" json:"user_id"`
	RightID int64 `db:"right_id" json:"right_id"`
}

func InitPersonsRepository(cnf *config.AppConfig, db *storage.DBConnector) *PersonsRepo {
	repo := &PersonsRepo{
		db:           db,
		systemRights: make(map[int64]string),
	}
	listData := make([]struct {
		RightID   int64  `db:"right_id"`
		RightName string `db:"right_name"`
	}, 0, 10)
	err := repo.db.Client.Select(&listData, "SELECT right_id, right_name FROM users_rights_names")
	if err != nil {
		logger.Fatal("error load system rights", err)
	}
	for i := range listData {
		repo.systemRights[listData[i].RightID] = listData[i].RightName
	}
	return repo
}

func (p *PersonsRepo) LoadPersons(offset int64) ([]Person, []PersonRight, error) {
	users := make([]Person, 0, 0)
	userRights := make([]PersonRight, 0, loadLimit*2)
	err := p.db.Client.Select(&users, "SELECT user_id, email, first_name, last_name, register_date, active FROM users ORDER BY user_id LIMIT ?, ?", offset, loadLimit)
	if err != nil {
		return users, userRights, err
	}
	userIds := make([]int64, 0, loadLimit)
	for i := range users {
		userIds = append(userIds, users[i].UserID)
	}

	query, args, errQ := sqlx.In("SELECT user_id, right_id FROM users_rights WHERE user_id IN (?) ORDER BY user_id", userIds)
	if errQ != nil {
		return users, userRights, errQ
	}

	err = p.db.Client.Select(&userRights, query, args...)
	if err != nil {
		return users, userRights, err
	}

	return users, userRights, nil
}

func (p *PersonsRepo) GetRightsMap() map[int64]string {
	return p.systemRights
}

func (p *PersonsRepo) UpdateUser(userID int64, firstName, lastName, email string, active bool) (bool, error) {
	if userID == 0 {
		return false, nil
	}
	prs := Person{
		Email:     email,
		LastName:  lastName,
		FirstName: firstName,
		UserID:    userID,
	}

	if active {
		prs.Active = true
	}
	res, err := p.db.Client.NamedExec("UPDATE users SET email=:email, first_name=:first_name, last_name=:last_name WHERE user_id=:user_id", prs)
	if err != nil {
		return false, err
	}
	count, errC := res.RowsAffected()
	if errC != nil {
		return false, errC
	}
	return count > 0, nil
}
