package main

import (
	"encrypt/colors"
	"encrypt/event"
	"encrypt/screens"
	"log"
)

func init() {
	screens.Clear()
	screens.Banner()
}

func main() {
	var ans string
	for ans != "close" {
		screens.Input(&ans)
		if event.Listen(ans) {
		}
		val, err := event.Read("./test.enc", ans)
		if err != nil {
			colors.SetColorln("Invalid Passphrase!", colors.Red)
			continue
		}
		colors.SetColorln(val, colors.Green)
	}
	data := "Hello, World! New World! Lorem Ipsum!"
	passphrase := "secret"
	fileName := "test"
	f, err := event.Create(fileName)
	if err != nil {
		log.Println(err)
	}

	encryptedData, err := event.Encrypt([]byte(data), passphrase)
	if err != nil {
		log.Println(err)
	}

	_, err = event.Write(f, encryptedData)
	if err != nil {
		log.Println(err)
	}

	val, err := event.Read("./test.enc", passphrase)
	if err != nil {
		log.Println(err)
	}

	colors.SetColorln(val, colors.Green)
}
