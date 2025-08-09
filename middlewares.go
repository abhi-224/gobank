package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func withJwt(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating the requesting using jwt middleware")

		tokenString, err := extractBearer(r.Header)
		if err != nil {
			log.Println(err.Error())
			WriteJson(w, http.StatusUnauthorized, err.Error())
			return
		}
		log.Printf("the bearer token is %v", tokenString)

		token, err := validateJwt(tokenString)
		if err != nil {
			log.Println(err.Error())
			WriteJson(w, http.StatusUnauthorized, err.Error())
			return
		}
		log.Printf("the extracted token %+v", token)
		h(w, r)
	}
}
func extractBearer(header http.Header) (string, error) {
	auth := header.Get("Authorization")
	if auth == "" {
		return "", errors.New("authorization header is missing")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(auth, prefix) {
		return "", errors.New("authorization header does not start with Bearer")
	}

	token := strings.TrimSpace(auth[len(prefix):])
	if token == "" {
		return "", errors.New("bearer token is empty")
	}

	return token, nil
}
func validateJwt(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	return token, nil
}

func createJwt(a *Account) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	claims := &jwt.RegisteredClaims{
		Issuer:    "gobank:pineapple",
		Subject:   strings.ToLower(a.FirstName) + "_" + strings.ToLower(a.LastName) + "_" + fmt.Sprint(a.Id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Duration(1).Minutes()))),
	}
	// claims := &jwt.MapClaims{
	// 	"issuer":        "gobank:pineapple",
	// 	"subjet":        strings.ToLower(a.FirstName) + "_" + strings.ToLower(a.LastName) + "_" + fmt.Sprint(a.Id),
	// 	"expiresAt":     jwt.NewNumericDate(time.Now().Add(time.Duration(time.Duration(1).Minutes()))),
	// 	"accountNumber": a.Number,
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
