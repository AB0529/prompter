package prompter

import (
	"fmt"

	"github.com/gookit/color"
	"golang.org/x/crypto/ssh/terminal"
)

// PasswordSelector hides an input
func PasswordSelector(q *Question, err error) (string, error) {
	// Print the question
	if err != nil {
		fmt.Print(fmt.Sprintf("\n[%s]", color.White.Sprint(err.Error())) + color.Red.Sprint(" > "))
	} else {
		fmt.Print(color.Cyan.Sprint(q.Message) + "\n" + color.Red.Sprint("> "))
	}

	password, err := terminal.ReadPassword(0)
	if err == nil {
		pass := string(password)

		return pass, nil
	}

	return "", err
}
