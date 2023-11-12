package repository

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
)

type categoryDomainRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
	UpdateCategory(*model.CategoryUpdate, int) (*model.Category, helper.Error)
}

type categoryRepo struct{}

var CategoryRepo categoryDomainRepo = &categoryRepo{}

func (c *categoryRepo) CreateCategory(category *model.Category) (*model.Category, helper.Error) {
	db := database.GetDB()

	err := db.Create(&category).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return category, nil
}

func (c *categoryRepo) UpdateCategory(categoryUpdated *model.CategoryUpdate, categoryId int) (*model.Category, helper.Error) {
	db := database.GetDB()

	var category model.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&category).Updates(categoryUpdated)

	return &category, nil
}
