package event

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Read(filePath string, passphrase string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	fmt.Println("@@@@@@@@@@@")
	if err != nil {
		log.Println(err)
		return "error:12", err
	}

	value := string(Decrypt(data, passphrase))

	return value, nil
}
