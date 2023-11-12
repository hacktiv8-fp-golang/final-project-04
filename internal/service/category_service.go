package service

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/repository"

	"github.com/asaskevich/govalidator"
)

type categoryServiceRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
}

type categoryService struct{}

var CategoryService categoryServiceRepo = &categoryService{}

func (c *categoryService) CreateCategory(category *model.Category) (*model.Category, helper.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	categoryResponse, err := repository.CategoryRepo.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return categoryResponse, nil
}
