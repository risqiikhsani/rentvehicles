package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Generates a JWT token with valid user_id and user_role
func Test_generate_jwt_token_with_valid_user_id_and_user_role(t *testing.T) {
	// Initialize test data
	user_id := uint(1)
	user_role := "admin"

	// Invoke code under test
	tokenString, err := GenerateJWTToken(user_id, user_role)

	// Assert the result
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}
