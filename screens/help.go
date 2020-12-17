package screens

import "encrypt/colors"

var Help = map[string]string{
	"-help": "[ Gives available commands ]",
}

func ShowCommands() {
	for k, v := range Help {
		colors.SetColorln(k, " > ", v)
	}
}
