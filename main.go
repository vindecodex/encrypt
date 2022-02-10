package main

import (
	"log"

	"github.com/vindecodex/encrypt/colors"
	"github.com/vindecodex/encrypt/event"
	"github.com/vindecodex/encrypt/screens"
	"github.com/vindecodex/encrypt/service"
)

func init() {
	screens.Clear()
	screens.Banner()
}

func main() {
	var ans string
	for ans != "close" {
		screens.Input(&ans)
		if event.Listen(&ans) {
			continue
		}
	}
	data := "Hello, World! New World! Lorem Ipsum!"
	passphrase := "secret"
	fileName := "test"
	f, err := service.Create(fileName)
	if err != nil {
		log.Println(err)
	}

	encryptedData, err := service.Encrypt([]byte(data), passphrase)
	if err != nil {
		log.Println(err)
	}

	_, err = service.Write(f, encryptedData)
	if err != nil {
		log.Println(err)
	}

	val, err := service.Read("./test.enc", passphrase)
	if err != nil {
		log.Println(err)
	}

	colors.SetColorln(val, colors.Green)
}
