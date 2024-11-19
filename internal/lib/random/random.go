package random

import (
	"crypto/rand"
	"fmt"
)

func NewRandomString(l int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if l <= 0 {
		return "", fmt.Errorf("invalid length: %d", l)
	}

	result := make([]byte, l)

	charsetLength := byte(len(charset))
	for i := 0; i < l; i++ {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return "", fmt.Errorf("failed to generate random byte: %w", err)
		}
		result[i] = charset[randomByte[0]%charsetLength]
	}

	return string(result), nil
}
