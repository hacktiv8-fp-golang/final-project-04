package repository

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
)

type userDomainRepo interface {
	Register(*model.User) (*model.User, helper.Error)
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
