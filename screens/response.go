package screens

import (
	"fmt"

	"github.com/vindecodex/encrypt/colors"
)

func Input(input interface{}) {
	colors.SetColor("encrypt: ", colors.Blue)
	fmt.Scanln(input)
}
func Response(res string) {
	colors.SetColorln(res, colors.Blue)
}
