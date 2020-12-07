package event

import (
	"io/ioutil"
	"log"
)

func Read(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return "error:12", err
	}

	value := string(data)

	return value, nil
}
