package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type ShippingAddress struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"not null"`
	Username  string `json:"username" gorm:"not null"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address" gorm:"not null"`
	IsDefault uint   `json:"isDefault" gorm:"default:1"`
}
