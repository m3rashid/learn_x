package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("JWT_SEC"))

func GenerateJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(signingKey) // this is the signing key
	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("Server is ready to receive requests")
	HandleRoutes()
}
