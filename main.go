package main

import (
	"encrypt/event"
	"fmt"
	"log"
)

func main() {
	f, err := event.Create("hello")
	if err != nil {
		log.Println(err)
	}

	_, err = event.Write(f, "aaa")
	if err != nil {
		log.Println(err)
	}

	val, err := event.Read("./hello.enc")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(val)
}
