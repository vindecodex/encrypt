package main

import (
	"encrypt/colors"
	"encrypt/event"
	"log"
)

func init() {
	colors.Blue("GGGGGGGGGG                                             GGG    ")
	colors.Blue("GGG                                                    GGG    ")
	colors.Blue("GGG                                                    GGG    ")
	colors.Blue("GGGGGGG    GGGGGb.   .dGGGGb GGGdGGG GGG  GGG GGGGGb.  GGGGGG ")
	colors.Blue("GGG        GGG  GGb dGGP     GGGP    GGG  GGG GGG  GGb GGG    ")
	colors.Blue("GGG        GGG  GGG GGG      GGG     GGG  GGG GGG  GGG GGG    ")
	colors.Blue("GGG        GGG  GGG YGGb.    GGG     YGGb GGG GGG dGGP YGGb.  ")
	colors.Blue("GGGGGGGGGG GGG  GGG   YGGGGP GGG       YGGGGG GGGGGP     YGGG ")
	colors.Blue("                                          GGG GGG             ")
	colors.Blue("                                     YGb dGGP GGG             ")
	colors.Blue("                                       YGGP   GGG             ")
}

func main() {
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
