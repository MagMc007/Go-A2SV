package infrastructure

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"task_manager/domain"
)

func GenerateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	if err != nil {
		return "", err
	}
	
		return tokenString, nil
}