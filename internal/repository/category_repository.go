package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"fmt"
)

type categoryDomainRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
	GetAllCategories() ([]*model.Category, helper.Error)
	GetCategoryById(int) (*model.Category, helper.Error)
	UpdateCategory(*model.CategoryUpdate, int) (*model.Category, helper.Error)
	DeleteCategory(int) (helper.Error)
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

func (s *categoryRepo) GetAllCategories() ([]*model.Category, helper.Error) {
	db := database.GetDB()
	var categories []*model.Category

	err := db.Preload("Products").Find(&categories).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return categories, nil
}

func (c *categoryRepo) GetCategoryById(categoryId int) (*model.Category, helper.Error) {
	db := database.GetDB()

	var category model.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, helper.NotFound(fmt.Sprintf("Category with id %d not found", categoryId))
	}

	return &category, nil
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

func (s *categoryRepo) DeleteCategory(categoryId int) helper.Error {
	db := database.GetDB()
	var category model.Category

	err := db.Where("id = ?", categoryId).Delete(&category).Error

	if err != nil {
		return helper.ParseError(err)
	}

	return nil
}
