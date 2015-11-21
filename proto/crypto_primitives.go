package proto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)


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
	Check(err)

	ciphertext := make([]byte, aes.BlockSize + len(data))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	Check(err)

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext
}

func Decrypt(key, data []byte) []byte {
	if len(key) < 32 {
		key = Hash(key)[:32]
	}
	block, err := aes.NewCipher(key)
	Check(err)
	if len(data) < aes.BlockSize {
		Check(errors.New("ciphertext too short"))
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return data
}
