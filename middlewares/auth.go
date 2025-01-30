package middlewares

import (
	"github.com/SthiraPs/DeliveryApp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization token",
		})
		return
	}

	email, roleId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.Set("email", email)
	context.Set("roleId", roleId)
	context.Next()
}
