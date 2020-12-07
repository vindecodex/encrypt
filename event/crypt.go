package event

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// To create 32 character passphrase
func createHash(key string) string {
	mdFive := md5.New()
	mdFive.Write([]byte(key))
	return hex.EncodeToString(mdFive.Sum(nil))
}

// Encrypt data with our passphrase
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}
