package middleware

import (
	"OrderTracker/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context) {
	// This function will be used as a middleware to authenticate the user.
	// This function will be called from the routes.go file.

	// Authenticate the user here.
	// If the user is not authenticated, return a 401 status code.
	// If the user is authenticated, call the next handler.

	// Get the token from the request.
	token := context.GetHeader("Authorization")

	// Validate the token.
	userId, role, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Set the user ID and role in the context.
	context.Set("userId", userId)
	context.Set("role", role)

	// Call the next handler.
	context.Next()
}
