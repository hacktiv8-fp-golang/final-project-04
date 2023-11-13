package model

import "time"

type TransactionHistory struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Quantity int `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
	TotalPrice int `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
	ProductID int `json:"product_id"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User *User
}
