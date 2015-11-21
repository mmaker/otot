package proto

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("hello, world!")
	hashSum := Hash(data)
	if len(hashSum) <= 0 {
		t.Fail()
	}
}

func TestEncryption(t *testing.T) {
	data := []byte("hello, world!")
	key := []byte("secret key")
	cipherText := Encrypt(key, data)
	if len(cipherText) <= 0 {
		t.Fail()
	}

	plainText := Decrypt(key, cipherText)
	if len(plainText) <= 0 {
		t.Fail()
	}
	if bytes.Compare(plainText, data) != 0 {
		t.Errorf("'%s' != '%s'", plainText, data)
	}
}
