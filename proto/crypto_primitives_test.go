package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	data := []byte("hello, world!")
	hashSum := Hash(data)
	assert.True(t, len(hashSum) > 0, "No hash received!")
}

func TestEncryption(t *testing.T) {
	data := []byte("hello, world!")
	key := []byte("secret key")
	cipherText := Encrypt(key, data)
	assert.True(t, len(cipherText) > 0, "No ciphertext recived!")
	plainText := Decrypt(key, cipherText)
	assert.True(t, len(plainText) > 0, "No plaintext received!")
	assert.Equal(t, plainText, data)
}
