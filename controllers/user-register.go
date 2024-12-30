package controllers

import (
	"net/http"

	"github.com/deyweson/go-cryptography-api/database"
	"github.com/deyweson/go-cryptography-api/internal/crypto"
	"github.com/deyweson/go-cryptography-api/models"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	if err := models.ValidatorUserData(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	encryptCreditCard, err := crypto.Encrypt(user.CreditCardToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	encryptUserDocument, err := crypto.Encrypt(user.UserDocument)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	user.CreditCardToken = encryptCreditCard
	user.UserDocument = encryptUserDocument

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}
