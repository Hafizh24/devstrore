package repository

import (
	"fmt"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user model.User) error {
	var sqlStatement = `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		`

	_, err := ur.DB.Exec(sqlStatement, user.Username, user.Email, user.Password)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create : %w", err))
		return err
	}
	return nil
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {

	var user model.User
	var sqlStatement = `
		SELECT id,username,email,password
		FROM users
		WHERE email = $1
		LIMIT 1
		`
	err := ur.DB.QueryRowx(sqlStatement, email).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmail : %w", err))
		return user, err
	}

	return user, nil

}
func (ur *UserRepository) GetByID(id int) (model.User, error) {

	var user model.User
	var sqlStatement = `
		SELECT id,username,email,password
		FROM users
		WHERE id = $1
		LIMIT 1
		`
	err := ur.DB.QueryRowx(sqlStatement, id).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByID : %w", err))
		return user, err
	}

	return user, nil

}
