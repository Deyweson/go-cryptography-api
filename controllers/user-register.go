package controllers

import (
	"net/http"

	"github.com/deyweson/go-cryptography-api/models"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, user)
}
