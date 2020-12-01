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

const DefaultUserVersion = 0

type user struct {
	ID      int64  `db:"user_id"`
	Email   string `db:"email"`
	Pass    string `db:"pass"`
	Version int64  `db:"version"`
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

func (ur *UserRepository) RegisterUser(mail, password string) (registered int64, exist bool, err error) {
	// check mail already exist
	row := ur.db.Client.QueryRow("SELECT user_id FROM users WHERE email = ?", mail)
	var userID int
	err = row.Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		// something wrong
		return 0, false, err
	}

	userExist := err != nil && err == sql.ErrNoRows
	if !userExist {
		// already exist
		return 0, true, nil
	}

	var passHash string
	passHash, err = utils.GeneratePassword(ur.passConf, password)
	if err != nil {
		return 0, false, err
	}
	res, err := ur.db.Client.NamedExec("INSERT INTO users (email, pass, version) VALUES (:email, :pass, :version)", user{
		Email:   mail,
		Pass:    passHash,
		Version: DefaultUserVersion,
	})

	if err != nil {
		return 0, false, err
	}

	uID, err := res.LastInsertId()
	if err != nil {
		return 0, false, err
	}

	return uID, false, nil
}

func (ur *UserRepository) LoginUser(mail, password string) (userID int64, userVersion int64, err error) {
	var p user
	err = ur.db.Client.Get(&p, "SELECT * FROM users WHERE email = ?", mail)
	if err != nil {
		return
	}
	valid, err := utils.ComparePassword(password, p.Pass)
	if !valid {
		return 0, 0, nil
	}
	userVersion = p.Version
	userID = p.ID
	return
}

func (ur *UserRepository) CheckVersion(userID, version int64) (valid bool, err error) {
	var p user
	err = ur.db.Client.Get(&p, "SELECT * FROM users WHERE user_id = ?", userID)
	if err != nil {
		return
	}
	return p.Version == version, nil
}
