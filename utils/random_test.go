package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Generates a random string of length 20
func TestRandomStringLength(t *testing.T) {
	result := RandomString()
	assert.Equal(t, 20, len(result))
}

// Uses a charset of alphanumeric characters
func TestRandomStringCharset(t *testing.T) {
	result := RandomString()
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range result {
		assert.Contains(t, charset, string(char))
	}
}

// None
func TestRandomStringNone(t *testing.T) {
	result := RandomString()
	assert.NotEmpty(t, result)
}

func TestReturnsStringWithUUIDAndTimestamp(t *testing.T) {
	result := RandomStringUuid()
	parts := strings.Split(result, "_")
	assert.Equal(t, 2, len(parts))
	assert.Regexp(t, "^[a-zA-Z0-9-]+$", parts[0])
	assert.Regexp(t, "^[0-9]+$", parts[1])
}

func TestGeneratesRandomUUIDString(t *testing.T) {
	result := RandomStringUuid()
	parts := strings.Split(result, "_")
	assert.Regexp(t, "^[a-zA-Z0-9-]+$", parts[0])
}

func TestRetrievesCurrentTimestamp(t *testing.T) {
	result := RandomStringUuid()
	parts := strings.Split(result, "_")
	assert.Regexp(t, "^[0-9]+$", parts[1])
}
