package screens

import "encrypt/colors"

const Help = make(map[string]string{
	"-help": "[ Gives available commands ]",
})

func ShowCommands() {
	for k, v := range Help {
		colors.SetColorln(k, " > ", v)
	}
}
