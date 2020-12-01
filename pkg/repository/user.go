package repository

import (
	"database/sql"
	"grinder/pkg/config"
	"grinder/pkg/storage"
	"grinder/pkg/utils"
)

type UserRepository struct {
	db       *storage.DBConnector
	passConf *utils.PasswordConfig
}

type user struct {
	ID    string `db:"user_id"`
	Email string `db:"email"`
	Pass  string `db:"pass"`
}

func InitUserRepository(cnf *config.AppConfig, db *storage.DBConnector) *UserRepository {
	return &UserRepository{
		db: db,
		passConf: &utils.PasswordConfig{
			Time:    1,
			Memory:  64 * 1024,
			Threads: 4,
			KeyLen:  32,
		},
	}
}

func (ur *UserRepository) RegisterUser(mail, password string) (registered, exist bool, err error) {
	// check mail already exist
	row := ur.db.Client.QueryRow("SELECT user_id FROM users WHERE email = ?", mail)
	var userID int
	err = row.Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		// something wrong
		return false, false, err
	}

	userExist := err != nil && err == sql.ErrNoRows
	if !userExist {
		// already exist
		return false, true, nil
	}

	var passHash string
	passHash, err = utils.GeneratePassword(ur.passConf, password)
	if err != nil {
		return false, false, err
	}
	_, err = ur.db.Client.NamedExec("INSERT INTO client (email, pass) VALUES (:name, :secret)", user{
		Email: mail,
		Pass:  passHash,
	})

	if err != nil {
		return false, false, err
	}

	return true, false, nil
}
