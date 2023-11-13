package repository

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
)

type userDomainRepo interface {
	Register(*model.User) (*model.User, helper.Error)
	Login(*model.LoginCredential) (*model.User, helper.Error)
	GetUserById(int) (*model.User, helper.Error)
	UpdateBalance(*model.User, int) (int, helper.Error)
}

type userRepo struct{}

var UserRepo userDomainRepo = &userRepo{}

func (u *userRepo) Register(user *model.User) (*model.User, helper.Error) {
	db := database.GetDB()

	err := db.Create(&user).Error

	if err != nil {
		return nil, helper.InternalServerError("Something went wrong")
	}

	return user, nil
}

func (u *userRepo) Login(userLogin *model.LoginCredential) (*model.User, helper.Error) {
	db := database.GetDB()

	var user model.User

	err := db.Where("email = ?", userLogin.Email).First(&user).Error

	if err != nil {
		return nil, helper.Unauthorized("Invalid email/password")
	}

	return &user, nil
}

func (u *userRepo) GetUserById(userId int) (*model.User, helper.Error) {
	db := database.GetDB()
	var user model.User

	err := db.First(&user, userId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return &user, nil
}

func (u *userRepo) UpdateBalance(userWithUpdatedBalance *model.User, userId int) (int, helper.Error) {
	db := database.GetDB()
	var user model.User

	err := db.First(&user, userId).Error

	if err != nil {
		return 0, helper.ParseError(err)
	}

	db.Model(&user).Updates(userWithUpdatedBalance)

	return user.Balance, nil
}
