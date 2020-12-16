package event

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

// To create 32 character passphrase
func createHash(key string) string {
	mdFive := md5.New()
	mdFive.Write([]byte(key))
	return hex.EncodeToString(mdFive.Sum(nil))
}

// Encrypt data with our passphrase
func Encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// panic(err.Error())
		log.Println(err.Error())
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		// panic(err.Error())
		log.Println(err.Error())
		return nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}

// Decrypt data with our passphrase
func Decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		// panic(err.Error())
		log.Println(err.Error(), "@")
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// panic(err.Error())
		log.Println(err.Error())
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		// panic(err.Error())
		log.Println(err.Error())
		return nil, err
	}
	return plainText, nil
}
