package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")
var tokenExpiration = getTokenExpiration()

func GenerateToken(email string, firstName string, lastName string, phoneNo string, roleId uint, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"phoneNo":   phoneNo,
		"roleId":    float64(roleId),
		"userType":  userType,
		"exp":       tokenExpiration,
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (string, int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token type")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", 0, err
	}

	isValidToken := token.Valid
	if !isValidToken {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("invalid token claims")
	}

	//email := claims["email"].(string)
	email := claims["email"].(string)
	roleId := int64(claims["roleId"].(float64))
	return email, roleId, nil
}

func getTokenExpiration() time.Duration {
	exp := os.Getenv("JWT_EXPIRATION")
	if exp == "" {
		return time.Hour * 2 // Default expiration time
	}
	duration, err := time.ParseDuration(exp)
	if err != nil {
		return time.Hour * 2 // Fallback to default if parsing fails
	}
	return duration
}
