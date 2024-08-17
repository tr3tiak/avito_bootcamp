package pkg

import (
	"avito_bootcamp/internal/entity"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(user *entity.User) (string, error) {
	logrus.Info("generate token started")
	claims := jwt.MapClaims{
		"sub":  user.Name,
		"role": user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := GetSecretKey()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logrus.Error("token signed error", err)
		return "", err
	}
	logrus.Info("Generate token complete")
	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(GetSecretKey()), nil
	})
	if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response := tokenClaims["role"].(string)
		return response, nil
	}
	return "", err

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

func ComparePassword(hashedPassword string, password string) error {
	logrus.Info("compare password started")

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logrus.Error("compareHash", err)
		return err
	}
	logrus.Info("Compare password over")
	return nil
}
