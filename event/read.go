package event

import (
	"log"
	"os"
)

func Read(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data := []byte{}
	length, err := f.Read(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	val := string(data[:length])
	return val, nil
}
