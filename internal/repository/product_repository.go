package repository

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
)

type productDomainRepo interface {
	CreateProduct(*model.Product) (*model.Product, helper.Error)
}

type productRepo struct{}

var ProductRepo productDomainRepo = &productRepo{}

func (p *productRepo) CreateProduct(product *model.Product) (*model.Product, helper.Error) {
	db := database.GetDB()

	err := db.Create(&product).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return product, nil
}
