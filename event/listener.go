package event

import (
	"github.com/vindecodex/encrypt/colors"
	"github.com/vindecodex/encrypt/handler"
	"github.com/vindecodex/encrypt/screens"
)

func Listen(cmd *string) bool {
	switch *cmd {
	case "help":
		colors.SetColorln("Available Commands: ", colors.White)
		screens.ShowCommands()
		return true
	case "clear":
		screens.Clear()
		return true
	case "open":
		val, err := handler.OpenFile(cmd)
		if err != nil {
			return false
		}
		colors.SetColorln(val, colors.Yellow)
		return true
	default:
		colors.SetColor("Unknown Command", colors.Red)
		colors.SetColorln(": help [ for helpful commands ]", colors.White)
		return true
	}
}
