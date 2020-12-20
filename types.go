package prompter

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/eiannone/keyboard"
	"github.com/k0kubun/go-ansi"
)

func printQuestion(q string) {
	fmt.Printf("%s\n%s ", Title.Sprint(q), InputChar.Sprint(">"))
}

func printQuestionWithError(q string, err error) {
	fmt.Printf("%s (%s)\n%s ", Title.Sprint(q), ValidateError.Sprint(err.Error()), InputChar.Sprint(">"))
}

// Inputer will handle the input type and get the response
func Inputer(t *Input, scanner *bufio.Scanner) (string, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		printQuestionWithError(t.Message, err)
	} else {
		printQuestion(t.Message)
	}
	scanner.Scan()

	// Get response
	resp := scanner.Text()

	// Handle validators with response
	for _, val := range t.Validators {
		// Run the validator function
		err = val(resp)
		if err != nil {
			goto PrintQuestion
		}
	}

	return resp, nil
}

// Booleaner will handle the boolean type and get the response
func Booleaner(t *Boolean, scanner *bufio.Scanner) (bool, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		fmt.Printf(fmt.Sprintf("[%s] ", BooleanPrompt.Sprint("Y/N")))
		printQuestionWithError(t.Message, err)
	} else {
		fmt.Printf(fmt.Sprintf("[%s] ", BooleanPrompt.Sprint("Y/N")))
		printQuestion(t.Message)
	}
	scanner.Scan()

	// Get response
	resp := scanner.Text()

	// Handle validators with response
	for _, val := range t.Validators {
		// Run the validator function
		err = val(resp)
		if err != nil {
			goto PrintQuestion
		}
	}

	if resp == "" {
		return true, nil
	}

	switch string(resp[0]) {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		err = errors.New("Value unknown")
		goto PrintQuestion
	}

}

// Passworder will handle the password type and get the response
func Passworder(t *Password) (string, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		printQuestionWithError(t.Message, err)
	} else {
		printQuestion(t.Message)
	}

	// Get the password
	resp, err := speakeasy.FAsk(os.Stdout, "")

	if err != nil {
		return "", err
	}

	// Handle validators with response
	for _, val := range t.Validators {
		// Run the validator function
		err = val(resp)
		if err != nil {
			goto PrintQuestion
		}
	}

	return resp, nil
}

// Multiselecter will handle the multiselect type and get the response
func Multiselecter(t *Multiselect) (string, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		fmt.Printf("%s (%s)\n", Title.Sprint(t.Message), ValidateError.Sprint(err.Error()))
	} else {
		fmt.Printf("%s\n", Title.Sprint(t.Message))
	}

	// Default response to first option
	resp := t.Options[0]
	// Print options
	for _, ms := range t.Options {
		fmt.Println(InputChar.Sprint("> ") + MultiselectOptions.Sprint(ms))
	}

	// Open keyboard event to get arrow keys
	keyboard.Open()
	defer func() { keyboard.Close() }()

	l := len(t.Options)
	i := l - 1
	// Move cursor up one
	ansi.CursorUp(1)
loop:
	for {
		// Handle key presses
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Move arrow key up, and approitate answer
		switch key {
		case keyboard.KeyArrowUp:
			i--
			if i < 0 {
				i++
				continue
			}

			ansi.CursorUp(1)
			break
		// Move arrow key down, and approitate answer
		case keyboard.KeyArrowDown:
			i++
			if i > l-1 {
				i = l - 1
				continue
			}
			ansi.CursorDown(1)
			break
		case keyboard.KeyCtrlC:
			keyboard.Close()
			os.Exit(0)
		default:
			// Load the answer on any other keypress
			fmt.Print("\r")
			for j := l; j > i; j-- {
				fmt.Println()
			}
			resp = t.Options[i]
			fmt.Println(MultiselectAnswer.Sprint(resp))
			break loop
		}
	}

	// Handle validators with response
	for _, val := range t.Validators {
		// Run the validator function
		err = val(resp)
		if err != nil {
			goto PrintQuestion
		}
	}

	return resp, nil
}
