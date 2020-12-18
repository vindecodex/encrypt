package handler

import (
	"encrypt/colors"
	"encrypt/event"
	"fmt"
	"log"
)

func OpenFile(cmd *string) (string, error) {
	colors.SetColor("File: ", colors.Blue)
	fmt.Scanln(cmd)
	file := *cmd
	colors.SetColor("Passphrase: ", colors.Blue)
	passphrase := *cmd

	val, err := event.Read(file, passphrase)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return val, nil
}
