package prompter

import (
	"bufio"
	"fmt"
	"os"
)

// TODO:
// - Fix multiselect going out of bounds
// - Multiselect allow for ctrl + to exit

// Multiselect select a single value in a list of multiple
type Multiselect []string

// Password empty struct to make it known it's a password
type Password struct{}

// YesNo empty struct to make it knowns it's a YesNo prompt
type YesNo struct{}

// Validator validates function passed into it, MUST RETURN ERROR
type Validator func(ans interface{}) error

// Question structure for a question that'll be asked
type Question struct {
	Message   string      `json:"message" binding:"required"`
	Name      string      `json:"name" binding:"required"`
	Validator []Validator `json:"validator,omitempty"`
	Type      interface{} `json:"type,omitempty"`
}

// Prompt the promt which will ask the questions provided and get the answers
type Prompt struct {
	Questions []*Question
}

// Ask will actually ask the questions and get the answers
func Ask(p *Prompt, v interface{}) error {
	var err error
	var answer string
	scanner := bufio.NewScanner(os.Stdin)

	// TODO: Add validators for multiselect and yes no
	for _, q := range p.Questions {
		// Handle Multiselect
		switch q.Type.(type) {
		case Multiselect:
			// Print question
			fmt.Println(Title.Sprint(q.Message))
			// Print options
			for _, ms := range q.Type.(Multiselect) {
				fmt.Println(InputChar.Sprint("> ") + MultiselectOptions.Sprint(ms))
			}
			// Get the answer
			WriteAnswer(v, q.Name, Multiselector(q))
			continue
		case Password:
		getPasswordAnswer:
			answer, err = PasswordSelector(q, err)
			for _, val := range q.Validator {
				err = val(answer)

				// Handle validator
				if q.Validator != nil && err != nil {
					goto getPasswordAnswer
				}
			}

			WriteAnswer(v, q.Name, answer)
			fmt.Println()
			continue
		// Yes no
		case YesNo:
		GetYesNoAnswer:
			b, err := HandleYesNo(q, scanner, err)

			if err != nil {
				goto GetYesNoAnswer
			}
			WriteAnswer(v, q.Name, b)
			continue
		}

	getAnswer:
		// Print the question
		if err != nil {
			fmt.Print(Title.Sprint(q.Message) + "\n" + fmt.Sprintf("[%s]", ValidateError.Sprint(err.Error())) + InputChar.Sprint(" > "))
		} else {
			fmt.Print(Title.Sprint(q.Message) + "\n" + InputChar.Sprint("> "))
		}
		scanner.Scan()

		answer := scanner.Text()
		for _, val := range q.Validator {
			err = val(answer)

			// Handle validator
			if q.Validator != nil && err != nil {
				goto getAnswer
			}
		}

		// Load the answer
		WriteAnswer(v, q.Name, answer)

	}

	return nil
}
