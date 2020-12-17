package prompter

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen will clear a terminal screen for any platform
func ClearScreen() {
	// Windows
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Everything else
		fmt.Println("\033[2J")
	}
}
