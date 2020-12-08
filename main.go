package main

import (
	"encrypt/colors"
	"encrypt/event"
	"log"
)

func init() {
	colors.SetColor("GGGGGGGGGG                                             GGG    ", colors.Blue)
	colors.SetColor("GGG                                                    GGG    ", colors.Blue)
	colors.SetColor("GGG                                                    GGG    ", colors.Blue)
	colors.SetColor("GGGGGGG    GGGGGb.   .dGGGGb GGGdGGG GGG  GGG GGGGGb.  GGGGGG ", colors.Blue)
	colors.SetColor("GGG        GGG  GGb dGGP     GGGP    GGG  GGG GGG  GGb GGG    ", colors.Blue)
	colors.SetColor("GGG        GGG  GGG GGG      GGG     GGG  GGG GGG  GGG GGG    ", colors.Blue)
	colors.SetColor("GGG        GGG  GGG YGGb.    GGG     YGGb GGG GGG dGGP YGGb.  ", colors.Blue)
	colors.SetColor("GGGGGGGGGG GGG  GGG   YGGGGP GGG       YGGGGG GGGGGP     YGGG ", colors.Blue)
	colors.SetColor("                                          GGG GGG             ", colors.Blue)
	colors.SetColor("                                     YGb dGGP GGG             ", colors.Blue)
	colors.SetColor("                                       YGGP   GGG             ", colors.Blue)
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
