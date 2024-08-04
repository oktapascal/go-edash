package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go-rental/enums"
	"time"
)

type JwtParameters struct {
	Email string
	Role  enums.Role
}

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
func GenerateToken(parameters *JwtParameters) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": viper.GetString("APP_NAME"),
		"sub": parameters.Email,
		"aud": parameters.Role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token using the provided secret key.
//
// The function takes a string token as input and uses the jwt.Parse function to parse and verify the token.
// It uses a custom function as the key function to validate the token's signature using the JWT_SIGNATURE_KEY.
//
// If the token is successfully parsed and verified, the function returns nil.
// If there is an error during parsing or verification, the function returns the corresponding error.
//
// If the token is not valid (expired, malformed, etc.), the function returns an error with the message "invalid token".
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.GetString("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
