package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create File
	f, err := os.Create("sample.enc")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Writing File
	_, err = f.WriteString("absdlkfj \n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
