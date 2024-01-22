package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(userPass string, userInputPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(userInputPass))
	if err != nil {
		return err
	}
	return nil
}