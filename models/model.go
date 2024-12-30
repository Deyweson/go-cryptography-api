package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserDocument    string `json:"user_document" binding:"required"`
	CreditCardToken string `json:"credit_card_token" binding:"required"`
	Value           int    `json:"value" binding:"required"`
}

//
// min=10
