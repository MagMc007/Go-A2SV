package infrastructure

import 	(
	"errors"
	"golang.org/x/crypto/bcrypt"

	"task_manager/domain"
)

// function to hash pwd
func HashPwd(pwd []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// function to check pwd
func CheckPwd(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}