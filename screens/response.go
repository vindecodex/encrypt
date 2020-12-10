package screens

import (
	"encrypt/colors"
	"fmt"
)

func Input(input interface{}) {
	colors.SetColor("encrypt: ", colors.Blue)
	fmt.Scanln(input)
}
func Response(res string) {
	colors.SetColorln(res, colors.Blue)
}
