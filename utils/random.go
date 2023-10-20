package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// randomstring generates a random string
func RandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 20 // Set your desired length

	// Initialize a random source using the current time
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rnd.Intn(len(charset))]
	}

	return string(b)
}

func RandomStringUuid() string {
	randomString := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return randomString + "_" + strconv.FormatInt(timestamp, 10)
}
