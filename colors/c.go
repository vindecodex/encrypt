package colors

import (
	"fmt"
	"strconv"
)

const (
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
)

func SetColor(str string, c int) {
	color := "\x1b[" + strconv.Itoa(c) + "m"
	fmt.Print(color, str)
}

func SetColorLn(str string, c int) {
	color := "\x1b[" + strconv.Itoa(c) + "m"
	fmt.Println(color, str)
}
