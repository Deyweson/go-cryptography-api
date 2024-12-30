package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserDocument    string `json:"user_document" validate:"min=3,max=25, regexp=^[0-9]*$"`
	CreditCardToken string `json:"credit_card_token" validate:"nonzero, len=16, regexp=^[0-9]*$"`
	Value           int    `json:"value" validate:"nonzero"`
}

func ValidatorUserData(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
