package routes

import (
	"github.com/deyweson/go-cryptography-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/user", controllers.UserRegister)

	r.Run(":8080")
}
