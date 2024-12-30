package controllers

import (
	"net/http"
	"strconv"

	"github.com/deyweson/go-cryptography-api/database"
	"github.com/deyweson/go-cryptography-api/internal/crypto"
	"github.com/deyweson/go-cryptography-api/models"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	var user models.User

	database.DB.First(&user, idInt)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "User not found",
		})
		return
	}

	decryptCreditCard, err := crypto.Decrypt(user.CreditCardToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	decryptUserDocument, err := crypto.Decrypt(user.UserDocument)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	user.CreditCardToken = decryptCreditCard
	user.UserDocument = decryptUserDocument

	c.JSON(http.StatusOK, user)
}
