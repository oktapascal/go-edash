package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var SECRET_KEY = []byte(viper.GetString("JWT_SIGNATURE_KEY"))

// GenerateToken generates a JWT token for the given email address.
//
// The function takes an email string as input and uses the provided secret key to sign a JWT token.
// The token contains the email address in its claims and has an expiration time of 24 hours.
//
// The function uses the jwt.NewWithClaims function to create a new token with the specified claims.
// It then calls the token's SignedString method to generate the token string using the secret key.
//
// If the token is successfully generated, the function returns the token string and nil.
// If there is an error during token signing, the function returns an empty string and the corresponding error.
func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token using the provided secret key.
//
// The function takes a string token as input and uses the jwt.Parse function to parse and verify the token.
// It uses a custom function as the key function to validate the token's signature using the SECRET_KEY.
//
// If the token is successfully parsed and verified, the function returns nil.
// If there is an error during parsing or verification, the function returns the corresponding error.
//
// If the token is not valid (expired, malformed, etc.), the function returns an error with the message "invalid token".
func VerifyToken(token string) error {
	parse, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return err
	}

	if !parse.Valid {
		return errors.New("invalid token")
	}

	return nil
}
