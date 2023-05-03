package service

import (
	"errors"

	"github.com/hafizh24/devstore/internal/app/model"
	"github.com/hafizh24/devstore/internal/app/repository"
	"github.com/hafizh24/devstore/internal/app/schema"
	"github.com/hafizh24/devstore/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repo repository.IUserRepository
}

func NewRegistrationService(repo repository.IUserRepository) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (rs *RegistrationService) Register(req *schema.RegisterReq) error {

	existingUser, _ := rs.repo.GetByEmail(req.Email)
	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	password, _ := rs.hashPassword(req.Password)

	var insertData model.User
	insertData.Email = req.Email
	insertData.Password = password
	insertData.Username = req.Username

	err := rs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (rs *RegistrationService) hashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytePassword), nil

}
