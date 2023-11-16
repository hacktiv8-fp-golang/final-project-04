package model

import "time"

type Product struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Title string `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price int `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock int `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int `json:"category_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductCreate struct {
	Title string `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price int `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock int `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int `json:"category_id"`
}

type ProductUpdate struct {
	Title string `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price int `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock int `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int `json:"category_id"`
}
