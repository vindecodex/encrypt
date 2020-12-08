package screens

import (
	"os"
	"os/exec"
	"runtime"
)

const (
	linux   = "linux"
	windows = "windows"
)

func Clear() {
	syst := runtime.GOOS

	switch syst {
	case linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case windows:
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
