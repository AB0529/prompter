package prompter

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Multiselector handles getting a result from a multiselector type question
func Multiselector(q *Question) interface{} {
	// Default answer is the first option
	answer := q.Multiselect[0]

	// Open keyboard event to get arrow keys
	keyboard.Open()
	defer func() { keyboard.Close() }()

	i := len(q.Multiselect) - 1
	// Move cursor up one
	fmt.Print("\033[A")

loop:
	for {
		// Handle key presses
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		// Move arrow key up, and approitate answer
		case keyboard.KeyArrowUp:
			if i < 0 {
				continue
			}

			i--
			fmt.Print("\033[A")
			break
			// Move arrow key down, and approitate answer
		case keyboard.KeyArrowDown:
			if i > len(q.Multiselect)-1 {
				i--
				continue
			}
			i++
			fmt.Print("\033[B")
			break
		default:
			// Load the answer on any other keypress
			fmt.Print("\r")
			for j := len(q.Multiselect); j > i; j-- {
				fmt.Println()
			}
			answer = q.Multiselect[i]
			fmt.Println(Magenta(answer))
			break loop
		}
	}

	return answer
}
