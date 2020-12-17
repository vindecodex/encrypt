package screens

import "encrypt/colors"

var Help = map[string]string{
	"help":  "[ Gives available commands ]",
	"close": "[ End encrypt process ]",
}

func ShowCommands() {
	for k, v := range Help {
		text := k + " > " + v
		colors.SetColorln(text, colors.White)
	}
}
