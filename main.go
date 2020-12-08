package main

import (
	"encrypt/colors"
	"encrypt/event"
	"encrypt/screens"
	"fmt"
	"log"
)

func init() {
	screens.Clear()
	screens.Banner()
}

func main() {
	var cls string
	for cls != "close" {
		fmt.Print("encrypt")
		fmt.Scanln(&cls)
	}
	data := "Hello, World! New World! Lorem Ipsum!"
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

	val, err := event.Read("./test.enc", passphrase)
	if err != nil {
		log.Println(err)
	}

	colors.SetColor(val, colors.Green)
}
