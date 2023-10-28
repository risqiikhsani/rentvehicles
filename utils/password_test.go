package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCheckPassword(t *testing.T) {

	// Define test cases
	testCases := []struct {
		password       string
		hashedPassword string
		expectError    bool
	}{
		{"password123", hashPassword("password123"), false},  // Matched password
		{"wrongpassword", hashPassword("password123"), true}, // Non-matching password
	}

	// Initialize the require package with the testing.T instance
	// require := require.New(t)
	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.password, func(t *testing.T) {
			err := CheckPassword(tc.password, tc.hashedPassword)
			assert.Equal(tc.expectError, (err != nil), "they should be equal")
		})
	}
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
