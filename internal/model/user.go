package model

import "time"

type User struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	FullName string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role string `gorm:"not null" json:"role" valid:"matches(admin|customer)"`
	Balance int `gorm:"not null" json:"balance" valid:"range(0|100000000)"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	TransactionHistories []TransactionHistory
}

type LoginCredential struct {
	Email string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type BalanceUpdate struct {
	Balance int `gorm:"not null" json:"balance" valid:"range(0|100000000)"`
}
