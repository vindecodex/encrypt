package handler

import (
	"fmt"
	"log"

	"github.com/vindecodex/encrypt/colors"
	"github.com/vindecodex/encrypt/service"
)

func OpenFile(cmd *string) (string, error) {
	colors.SetColor("File: ", colors.Blue)
	fmt.Scanln(cmd)
	file := *cmd
	colors.SetColor("Passphrase: ", colors.Blue)
	fmt.Scanln(cmd)
	passphrase := *cmd

	val, err := service.Read(file, passphrase)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return val, nil
}
