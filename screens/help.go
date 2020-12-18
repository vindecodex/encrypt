package screens

import "encrypt/colors"

var Help = map[string]string{
	"help":  "[ Gives available commands ]",
	"close": "[ End encrypt process ]",
	"clear": "[ Clear the terminal ]",
}

func ShowCommands() {
	for k, v := range Help {
		text := k + " > " + v
		colors.SetColorln(text, colors.White)
	}
}
