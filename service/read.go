package service

import (
	"io/ioutil"
	"log"
)

func Read(filePath string, passphrase string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return "error:12", err
	}

	value, err := Decrypt(data, passphrase)
	if err != nil {
		log.Println(err)
		return "error:18", err
	}

	decryptedValue := string(value)

	return decryptedValue, nil
}
