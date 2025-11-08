package Utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/ticket/Config"
)

type JWTClaims struct {
	UserId          int       `json:"user_id"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Expiration_date time.Time `json:"expiration_date"`
	jwt.RegisteredClaims
}

func CreateToken(user_id uint, name string, phone string, email string, remember_me bool) (string, time.Time, error) {
	months := 1
	if remember_me {
		months = 6
	}

	expire_time := time.Now().AddDate(0, months, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         user_id,
		"name":            name,
		"phone":           phone,
		"email":           email,
		"expiration_date": expire_time,
	})
	tokenString, err := token.SignedString([]byte(Config.JWT_KEY))
	if err != nil {
		return "", time.Time{}, err
	}
	return string(tokenString), expire_time, nil
}

func VerifyToken(tokenString string) (uint, string, string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.JWT_KEY), nil
	})

	if err != nil {
		return 0, "", "", "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return 0, "", "", "", fmt.Errorf("invalid token")
	}

	if claims.Expiration_date.Unix() < time.Now().Unix() {
		return 0, "", "", "", fmt.Errorf("expired token")
	}

	return uint(claims.UserId), claims.Name, claims.Email, claims.Phone, nil
}
