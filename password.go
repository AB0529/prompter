package prompter

import (
	"fmt"
	"os"

	"github.com/bgentry/speakeasy"
)

// PasswordSelector hides an input
func PasswordSelector(q *Question, err error) (string, error) {
	// Print the question
	if err != nil {
		fmt.Print(fmt.Sprintf("\n[%s]", ValidateError.Sprint(err.Error())) + InputChar.Sprint(" > "))
	} else {
		fmt.Print(Title.Sprint(q.Message) + "\n" + InputChar.Sprint("> "))
	}

	pass, err := speakeasy.FAsk(os.Stdout, "")

	if err != nil {
		return "", err
	}

	return pass, nil
}
