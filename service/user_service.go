package service

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"kasir-app/dto"
	"kasir-app/errors"
	"kasir-app/model"
	"kasir-app/repository"
)

type IUserService interface {
	CreateUser(user *dto.CreateUserReq) error
}

type UserServiceImpl struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) IUserService {
	return &UserServiceImpl{repo: repo}
}

func (service *UserServiceImpl) CreateUser(userReq *dto.CreateUserReq) error {
	encryptedPass, err := bcrypt.GenerateFromPassword(
		[]byte(userReq.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		fmt.Println("errors encrypting password")
		return errors.Internal(err.Error())
	}

	userDb := model.User{
		Id:       uuid.New().String(),
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(encryptedPass),
	}

	err = service.repo.Create(&userDb)
	if err != nil {
		fmt.Println("errors creating user")
		return errors.BadRequest(err.Error())
	}
	return nil
}
