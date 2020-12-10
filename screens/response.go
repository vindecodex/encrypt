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
	colors.SetColorLn(res, colors.Blue)
}
