package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// This function will hash a password.
	// This function will be called from the controllers.go file.

	// Hash the password here.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// handle error appropriately, for now, we will return an empty string
		//Error() converts the error to a string
		return err.Error()
	}
	password = string(hashedPassword)

	return password
}

func ComparePassword(hashedPassword, password string) bool {
	// This function will compare a hashed password with a password.
	// This function will be called from the controllers.go file.

	// Compare the password here.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// handle error appropriately, for now, we will return false if there is an error
	return err == nil
}
