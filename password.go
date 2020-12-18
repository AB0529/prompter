package prompter

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

// PasswordSelector hides an input
func PasswordSelector(q *Question, err error) (string, error) {
	// Print the question
	if err != nil {
		fmt.Print(fmt.Sprintf("\n[%s]", ValidateError.Sprint(err.Error())) + InputChar.Sprint(" > "))
	} else {
		fmt.Print(Title.Sprint(q.Message) + "\n" + InputChar.Sprint("> "))
	}

	password, err := terminal.ReadPassword(0)
	if err == nil {
		pass := string(password)

		return pass, nil
	}

	return "", err
}
