package middlewares

import (
	"net/http"

	"github.com/RevanthGovindan/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	var token = context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
