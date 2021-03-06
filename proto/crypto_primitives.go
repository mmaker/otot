package proto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"log"
)

const BITS = 512

func check(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func Encrypt(key, data []byte) []byte {
	if len(key) < 32 {
		key = Hash(key)[:32]
	}
	block, err := aes.NewCipher(key)
	check(err)

	ciphertext := make([]byte, aes.BlockSize + len(data))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	check(err)

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext
}

func Decrypt(key, data []byte) []byte {
	if len(key) < 32 {
		key = Hash(key)[:32]
	}
	block, err := aes.NewCipher(key)
	check(err)
	if len(data) < aes.BlockSize {
		check(errors.New("ciphertext too short"))
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return data
}
