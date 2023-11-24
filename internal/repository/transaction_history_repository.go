package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
)

type transactionDomainRepo interface {
	CreateTransaction(
		user *model.User,
		product *model.Product,
		transaction *model.TransactionCreate,
		totalPrice int) (*model.TransactionHistory, helper.Error)

		GetTransactionsByUserID (userID uint)([]*model.TransactionHistory, helper.Error)
		GetAllTransaction ()([]model.TransactionHistory, helper.Error)
}

type transactionRepo struct{}

var TransactionRepo transactionDomainRepo = &transactionRepo{}

func (t *transactionRepo) CreateTransaction(
	user *model.User,
	product *model.Product,
	transaction *model.TransactionCreate,
	totalPrice int) (*model.TransactionHistory, helper.Error) {
		db := database.GetDB()

		tx := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		product.Stock -= transaction.Quantity
		user.Balance -= totalPrice

		if err := tx.Save(product).Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to update product stock")
		}

		if err := tx.Save(user).Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to update user balance")
		}

		category := &model.Category{} // Inisialisasi objek kategori
		if err := tx.Where("id = ?", product.CategoryID).First(category).Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to find category")
		}

		category.SoldProductAmount += transaction.Quantity

		if err := tx.Save(category).Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to update category sold product amount")
		}

		transactionHistory := &model.TransactionHistory{
			Quantity:   transaction.Quantity,
    	TotalPrice: totalPrice,
    	UserID:     user.ID,
    	ProductID:  product.ID,
			Product: product,
		}

		if err := tx.Create(transactionHistory).Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to create new transaction: " + err.Error())
		}

		if err := tx.Model(user).Association("TransactionHistories").Append(transactionHistory); err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to update user's transaction histories: ")
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return nil, helper.InternalServerError("Failed to commit transaction")
		}

		return transactionHistory, nil
	}


	func (t *transactionRepo) GetTransactionsByUserID (userID uint)([]*model.TransactionHistory, helper.Error){
		db := database.GetDB()

		var transactionUser []*model.TransactionHistory

		err := db.Preload("Product").Where("user_id = ?", userID).Find(&transactionUser).Error

		if err != nil {
			return nil, helper.ParseError(err)
		}

		return transactionUser, nil
	}


	func (t *transactionRepo) GetAllTransaction ()([]model.TransactionHistory, helper.Error){
		db := database.GetDB()

		var AllTransaction []model.TransactionHistory

		err := db.Preload("Product").Preload("User").Find(&AllTransaction).Error

		if err != nil {
			return nil, helper.ParseError(err)
		}
		return AllTransaction, nil
	}

