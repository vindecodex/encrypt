package event

import "fmt"

func Listen(cmd *string) bool {

	fmt.Println(*cmd)
	return true
	// switch cmd {
	// case "help":
	// colors.SetColorln("Available Commands: ", colors.White)
	// screens.ShowCommands()
	// return true
	// case "clear":
	// screens.Clear()
	// return true
	// case "open":
	// handler.OpenFile()
	// return true
	// default:
	// colors.SetColor("Unknown Command", colors.Red)
	// colors.SetColorln(": help [ for helpful commands ]", colors.White)
	// return true
	// }
}
