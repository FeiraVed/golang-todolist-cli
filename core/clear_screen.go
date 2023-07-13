package core

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	text := ""
	Os := runtime.GOOS
	if Os == "linux" {
		text = "clear"
	} else if Os == "android" {
		text = "clear"
	} else if Os == "windows" {
		text = "cls"
	}

	cmd := exec.Command(text) //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
