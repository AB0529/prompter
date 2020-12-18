package prompter

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/gookit/color"
)

// Multiselector handles getting a result from a multiselector type question
func Multiselector(q *Question) interface{} {
	// Default answer is the first option
	answer := q.Type.(Multiselect)[0]

	// Open keyboard event to get arrow keys
	keyboard.Open()
	defer func() { keyboard.Close() }()

	i := len(q.Type.(Multiselect)) - 1
	// Move cursor up one
	CursorUp(1)

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
			i--
			if i < 0 {
				i++
				continue
			}

			CursorUp(1)
			break
			// Move arrow key down, and approitate answer
		case keyboard.KeyArrowDown:
			i++
			if i > len(q.Type.(Multiselect))-1 {
				i = len(q.Type.(Multiselect)) - 1
				continue
			}
			CursorDown(1)
			break
		default:
			// Load the answer on any other keypress
			fmt.Print("\r")
			for j := len(q.Type.(Multiselect)); j > i; j-- {
				fmt.Println()
			}
			answer = q.Type.(Multiselect)[i]
			fmt.Println(color.Cyan.Sprint(answer))
			break loop
		}
	}

	return answer
}
