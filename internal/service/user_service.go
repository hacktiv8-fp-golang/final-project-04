package service

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/repository"

	"github.com/asaskevich/govalidator"
)

type userServiceRepo interface {
	Register(*model.User) (*model.User, helper.Error)
	Login(*model.LoginCredential) (string, helper.Error)
}

type userService struct{}

var UserService userServiceRepo = &userService{}

func (u *userService) Register(user *model.User) (*model.User, helper.Error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	password, err := helper.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = password

	userResponse, err := repository.UserRepo.Register(user)

	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *userService) Login(userLogin *model.LoginCredential) (string, helper.Error) {
	if _, err := govalidator.ValidateStruct(userLogin); err != nil {
		return "", helper.BadRequest(err.Error())
	}

	user, err := repository.UserRepo.Login(userLogin)

	if err != nil {
		return "", err
	}

	if isPasswordCorrect := helper.ComparePassword(userLogin.Password, user.Password); !isPasswordCorrect {
		return "", helper.Unauthorized("Invalid email/password")
	}

	token, err := helper.GenerateToken(user.ID, user.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}
