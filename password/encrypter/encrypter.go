package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("Нет параметра KEY")
	}
	return &Encrypter{
		Key: key,
	}
}

func (e *Encrypter) Encrypt(data []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		log.Fatal(err)
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Fatal(err)
	}
	return aesGSM.Seal(nonce, nonce, data, nil)
}

func (e *Encrypter) Decrypt(encryptedData []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		log.Fatal(err)
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	nonceSize := aesGSM.NonceSize()
	nonce := encryptedData[:nonceSize]
	cipherText := encryptedData[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatal(err)
	}
	return plainText
}
