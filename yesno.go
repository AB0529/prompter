package prompter

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// HandleYesNo handles the yes no prompt
func HandleYesNo(q *Question, scanner *bufio.Scanner, err error) (bool, error) {
	// Print the question
	fmt.Print(Yellow.Sprint("[Y/N] ") + Title.Sprint(q.Message) + "\n" + InputChar.Sprint("> "))

	scanner.Scan()
	t := strings.ToLower(scanner.Text())
	if t == "" {
		return true, nil
	}

	switch string(t[0]) {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		return false, errors.New("Unknown value for bool")
	}

}
