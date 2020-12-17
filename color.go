package prompter

import "fmt"

var (
	// Black color black
	Black = Color("\033[1;30m%s\033[0m")
	// Red color red
	Red = Color("\033[1;31m%s\033[0m")
	// Green color green
	Green = Color("\033[1;32m%s\033[0m")
	// Yellow color yellow
	Yellow = Color("\033[1;33m%s\033[0m")
	// Purple color purple
	Purple = Color("\033[1;34m%s\033[0m")
	// Magenta color magenta
	Magenta = Color("\033[1;35m%s\033[0m")
	// Teal color teal
	Teal = Color("\033[1;36m%s\033[0m")
	// White color white
	White = Color("\033[1;37m%s\033[0m")
)

// Color colorizes strings
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
