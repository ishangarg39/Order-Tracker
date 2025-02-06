package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("MySecretKey")

func GetTokens(email string, userId int64, role string) (string, error) {
	// This function will get the tokens from the request.
	// This function will be called from the controllers.go file.

	// Get the tokens here.
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (int64, string, error) {

	// This function will validate the token.
	// This function will be called from the controllers.go file.

	// Validate the token here.
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//Return the key for validation
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	//Extract the claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, "", err
	}

	//Returns the user ID by type asserting the value to float64
	userId := int64(claims["userId"].(float64))
	role := claims["role"].(string)
	return userId, role, nil
}
