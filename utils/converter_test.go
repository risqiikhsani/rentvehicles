package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should convert a valid string to uint without error
func TestConvertToUintValidString(t *testing.T) {
	str := "123"
	expected := uint(123)

	result, err := ConvertToUint(str)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestConvertToUintCorrectValue(t *testing.T) {
	str := "456"
	expected := uint(456)

	result, err := ConvertToUint(str)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestConvertToUintEmptyString(t *testing.T) {
	str := ""
	expectedErr := errors.New("strconv.ParseUint: parsing \"\": invalid syntax")

	result, err := ConvertToUint(str)

	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, uint(0), result)
}
