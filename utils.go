// +build !windows

package prompter

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearScreen will clear a terminal screen for any platform
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CursorUp moves cursor up n cells
func CursorUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

// CursorDown moves cursor down n cells
func CursorDown(n int) {
	fmt.Printf("\x1b[%dB", n)
}
