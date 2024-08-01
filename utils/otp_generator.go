package utils

import (
	"crypto/rand"
	"math/big"
)

// OTPGenerator generates a random OTP (One-Time Password) of the specified length.
// It returns a string containing the generated OTP and an error if any occurred.
//
// Parameters:
// - length: the desired length of the OTP.
//
// Returns:
// - string: the generated OTP.
// - error: an error if any occurred.
func OTPGenerator(length int) (string, error) {
	// Define the character set for generating the OTP.
	const charset = "0123456789"

	// Create a byte slice of the specified length.
	result := make([]byte, length)

	// Iterate over the indices of the result byte slice.
	for i := range result {
		// Generate a random number within the range of the character set length.
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// Return an empty string and the error if any occurred.
			return "", err
		}

		// Assign the character at the generated index to the current index of the result byte slice.
		result[i] = charset[num.Int64()]
	}

	// Convert the result byte slice to a string and return it.
	return string(result), nil
}
