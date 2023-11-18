package service

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/repository"

	"github.com/asaskevich/govalidator"
)

type transactionServiceRepo interface {
	CreateTransaction(transaction *model.TransactionCreate, userId int) (*model.TransactionHistory, helper.Error)
	GetTransactionsByUserID (userID uint)([]*model.TransactionHistory, helper.Error)
	GetAllTransaction () ([]model.TransactionHistory, helper.Error)
}

type transactionService struct{}

var TransactionService transactionServiceRepo = &transactionService{}

func (t *transactionService) CreateTransaction(transaction *model.TransactionCreate, userId int) (*model.TransactionHistory, helper.Error) {

	if _, err := govalidator.ValidateStruct(transaction); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	user, err := repository.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	product, err := repository.ProductRepo.GetProductById(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Stock < transaction.Quantity {
		return nil, helper.BadRequest("Insufficient product stock")
	}

	totalPrice := transaction.Quantity * product.Price

	if user.Balance < totalPrice {
		return nil, helper.BadRequest("Insufficient user balance")
	}


	createdTransaction, err := repository.TransactionRepo.CreateTransaction(user, product, transaction, totalPrice)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

func (t *transactionService) GetTransactionsByUserID (userID uint)([]*model.TransactionHistory, helper.Error){
	transactionResponse, err := repository.TransactionRepo.GetTransactionsByUserID(userID)

	if err != nil {
		return nil, err
	}
	
	return transactionResponse, nil
}

func (t *transactionService) GetAllTransaction () ([]model.TransactionHistory, helper.Error){
	transactionResponse, err := repository.TransactionRepo.GetAllTransaction()

	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}
