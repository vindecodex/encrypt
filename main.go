package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("sample.enc")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString("absdlkfj \n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
