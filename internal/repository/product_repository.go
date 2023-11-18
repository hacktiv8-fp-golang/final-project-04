package repository

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
)

type productDomainRepo interface {
	CreateProduct(*model.Product) (*model.Product, helper.Error)
	GetAllProducts() ([]*model.Product, helper.Error)
	UpdateProduct(*model.ProductUpdate, int) (*model.Product, helper.Error)
	DeleteProduct(int) (helper.Error)
	GetProductById(int) (*model.Product, helper.Error)
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

func (s *productRepo) GetAllProducts() ([]*model.Product, helper.Error) {
	db := database.GetDB()
	var products []*model.Product

	err := db.Find(&products).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return products, nil
}


func (c *productRepo) UpdateProduct(update *model.ProductUpdate, productId int) (*model.Product, helper.Error) {
	db := database.GetDB()

	var product model.Product

	err := db.First(&product, productId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&product).Updates(update)

	return &product, nil
}

func (s *productRepo) DeleteProduct(productId int) helper.Error {
	db := database.GetDB()
	var product model.Product

	err := db.Where("id = ?", productId).Delete(&product).Error

	if err != nil {
		return helper.ParseError(err)
	}

	return nil
}

func (s *productRepo) GetProductById(productId int) (*model.Product, helper.Error){
	db := database.GetDB()
	var product model.Product

	err := db.First(&product, productId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return &product, nil
}