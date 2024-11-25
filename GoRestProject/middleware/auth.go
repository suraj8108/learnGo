package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/utils"
)

func Authenticate(context *gin.Context) {
	//Validate JWT Token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// This Abort will not execute any further handler
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token not valid", "errorMsg": err.Error()})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
