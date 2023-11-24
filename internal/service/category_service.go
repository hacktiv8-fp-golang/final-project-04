package service

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/repository"

	"github.com/asaskevich/govalidator"
)

type categoryServiceRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
	GetAllCategories() ([]*model.Category, helper.Error)
	UpdateCategory(*model.CategoryUpdate, int) (*model.Category, helper.Error)
	DeleteCategory(int) (helper.Error)
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

func (s *categoryService) GetAllCategories() ([]*model.Category, helper.Error) {
	categories, err := repository.CategoryRepo.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryService) UpdateCategory(category *model.CategoryUpdate, categoryId int) (*model.Category, helper.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	categoryResponse, err := repository.CategoryRepo.UpdateCategory(category, categoryId)

	if err != nil {
		return nil, err
	}

	return categoryResponse, nil
}

func (s *categoryService) DeleteCategory(categoryId int) helper.Error {
	err := repository.CategoryRepo.DeleteCategory(categoryId)

	if err != nil {
		return err
	}

	return nil
}
