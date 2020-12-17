package event

import "encrypt/screens"

func Listen(cmd string) bool {
	switch cmd {
	case "clear":
		screens.Clear()
		return true
	default:
		return true
	}
}
