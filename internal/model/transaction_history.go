package model

import (
	"time"
)

type TransactionHistory struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Quantity int `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
	TotalPrice int `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
	ProductID int `json:"product_id"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product *Product
	User *User
}

type TransactionCreate struct {
	ProductID int `json:"product_id"`
	Quantity int `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
}

type TransactionCreateResponse struct {
	Message string `json:"message"`

	// TransactionBill represents the transaction bill information.
	TransactionBill struct {
		TotalPrice     int    `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
		Quantity       int    `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
		ProductTitle   string `json:"product_title"`
	} `json:"transaction_bill"`
}

type GetTransactionUserResponse struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Quantity int `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
	TotalPrice int `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
	ProductID int `json:"product_id"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product *Product
}