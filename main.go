package main

import (
	"encrypt/event"
	"fmt"
	"log"
)

func main() {
	data := "Hello, World!"
	passphrase := "secret"
	fileName := "test"
	f, err := event.Create(fileName)
	if err != nil {
		log.Println(err)
	}

	_, err = event.Write(f, event.Encrypt([]byte(data), passphrase))
	if err != nil {
		log.Println(err)
	}

	val, err := event.Read("./hello.enc")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(val)
}
