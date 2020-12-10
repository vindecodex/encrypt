package screens

import (
	"encrypt/colors"
)

func Input(input interface{}) {
	colors.SetColor("encrypt: ", colors.Blue)
}
func Response(res string) {
	colors.SetColorLn(res, colors.Blue)
}
