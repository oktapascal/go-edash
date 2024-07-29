package utils

import "golang.org/x/crypto/bcrypt"

// Hash takes a plaintext string and returns its bcrypt hashed representation and an error if any.
// The cost parameter is set to 14, which is a reasonable trade-off between security and performance.
//
// Parameters:
// - value: the plaintext string to be hashed
//
// Returns:
// - string: the bcrypt hashed representation of the plaintext string
// - error: an error if any occurred during the hashing process
func Hash(value string) (string, error) {
	// Generate the bcrypt hash of the plaintext string
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)

	// Return the hashed string and any error that occurred
	return string(bytes), err
}

// CheckHash compares a plaintext string with a bcrypt hashed string.
// It returns true if the plaintext matches the hash, false otherwise.
//
// Parameters:
// - value: the plaintext string to compare
// - hash: the bcrypt hashed string to compare against
//
// Returns:
// - bool: true if the plaintext matches the hash, false otherwise
func CheckHash(value string, hash string) bool {

	// Convert the hash and value to byte slices
	hashBytes := []byte(hash)
	valueBytes := []byte(value)

	// Compare the hash and value using bcrypt
	err := bcrypt.CompareHashAndPassword(hashBytes, valueBytes)

	// Return true if there is no error, false otherwise
	return err == nil
}
