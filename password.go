package prompter

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

// PasswordSelector hides an input
func PasswordSelector(q *Question) string {
	// Print question
	fmt.Print(Purple(q.Message) + "\n" + Red("> "))

	password, err := terminal.ReadPassword(0)
	if err == nil {
		return string(password)
	}

	return ""
}
