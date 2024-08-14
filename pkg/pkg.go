package pkg

import (
	"avito_bootcamp/internal/entity"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.Name,
		"role": user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	secretKey := GetSecretKey()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return GetSecretKey(), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return err

}

func GetSecretKey() string {
	return os.Getenv("SECRET_KEY")
}

func EncryptedPassword(pass string) (string, error) {
	password := []byte(pass)
	encPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(encPass), nil

}

func ComparePassword(EncryptedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(EncryptedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
