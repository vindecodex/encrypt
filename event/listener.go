package event

import "fmt"

func Listen(cmd string) bool {
	fmt.Println(cmd)
	return true
}
