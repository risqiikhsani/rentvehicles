package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckPassword(t *testing.T) {
	// Define test cases
	testCases := []struct {
		password       string
		hashedPassword string
		expectError    bool
	}{
		{"password123", "$2a$14$CZ4DTiN5r4..Qq3Je0TRcuwLOjC/q9Yzm7zNZQdZ54vPXrOeP85LS", false},  // Matched password
		{"wrongpassword", "$2a$14$CZ4DTiN5r4..Qq3Je0TRcuwLOjC/q9Yzm7zNZQdZ54vPXrOeP85LS", true}, // Non-matching password
	}

	// Initialize the require package with the testing.T instance
	require := require.New(t)

	for _, tc := range testCases {
		t.Run(tc.password, func(t *testing.T) {
			err := CheckPassword(tc.password, tc.hashedPassword)
			if tc.expectError {
				require.Error(err, "Expected an error for password: %s", tc.password)
			} else {
				require.NoError(err, "Expected no error for password: %s", tc.password)
			}
		})
	}
}
