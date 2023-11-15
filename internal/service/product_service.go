package service

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/repository"

	"github.com/asaskevich/govalidator"
)

type productServiceRepo interface {
	CreateProduct(*model.Product) (*model.Product, helper.Error)
}

type productService struct{}

var ProductService productServiceRepo = &productService{}

func (p *productService) CreateProduct(product *model.Product) (*model.Product, helper.Error) {
	if _, err := govalidator.ValidateStruct(product); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	if product.Stock < 5 {
		return nil, helper.BadRequest("The product stock must not be less than 5")
	}

	if _, err := repository.CategoryRepo.GetCategoryById(product.CategoryID); err != nil {
		return nil, err
	}

	productResponse, err := repository.ProductRepo.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return productResponse, nil
}
