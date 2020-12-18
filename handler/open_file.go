package handler

import (
	"encrypt/colors"
	"fmt"
)

func OpenFile(cmd *string) {
	colors.SetColor("File: ", colors.Blue)
	fmt.Scanln(cmd)
	fmt.Println(*cmd)
}
