package server

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) GenerateHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	return string(hash), err
}

func (user *User) CheckPassword(hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
}

func (user *User) GenerateToken() (string, error) {
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsTokenValid(authToken string) bool {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an token validation error")
		}
		return []byte("secret"), nil
	})
	return err == nil && token.Valid
}
