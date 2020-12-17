package event

import (
	"encrypt/colors"
	"encrypt/screens"
)

func Listen(cmd string) bool {
	switch cmd {
	case "clear":
		screens.Clear()
		return true
	default:
		colors.SetColorln("Unknown Command: -help [ for helpful commands ]", colors.White)
		return true
	}
}
