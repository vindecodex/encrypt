package screens

import "encrypt/colors"

type CMD string

var Help = map[CMD]string{
	"-help": "[ Gives available commands ]",
}

func ShowCommands() {
	for k, v := range Help {
		text := string(k + " > " + v)
		colors.SetColorln(text, colors.White)
	}
}
