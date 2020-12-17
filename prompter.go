package prompter

import (
	"bufio"
	"fmt"
	"os"
)

// Multiselect select a single value in a list of multiple
type Multiselect []string

// Validator validates function passed into it, MUST RETURN ERROR
type Validator func(ans interface{}) error

// Question structure for a question that'll be asked
type Question struct {
	Message     string      `json:"message" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Validator   Validator   `json:"validator,omitempty"`
	Multiselect Multiselect `json:"multiselect,omitempty"`
}

// Prompt the promt which will ask the questions provided and get the answers
type Prompt struct {
	Questions []*Question
}

// Ask will actually ask the questions and get the answers
func Ask(p *Prompt, v interface{}) error {
	scanner := bufio.NewScanner(os.Stdin)

	for _, q := range p.Questions {
		// Handle Multiselect
		if q.Multiselect != nil {
			// Print question
			fmt.Println(Purple(q.Message))
			// Print options
			for _, ms := range q.Multiselect {
				fmt.Println(Red("> ") + Yellow(ms))
			}
			// Get the answer
			WriteAnswer(v, q.Name, Multiselector(q))
			continue
		}
		// TODO: better error handling
		var err error
	getAnswer:
		answer := func(err error) string {
			// Print the question
			if err != nil {
				fmt.Print(Purple(q.Message) + "\n" + fmt.Sprintf("[%s]", White(err.Error())) + Red(" > "))
			} else {
				fmt.Print(Purple(q.Message) + "\n" + Red("> "))
			}
			scanner.Scan()

			return scanner.Text()
		}(err)
		err = q.Validator(answer)

		// Handle validator
		if q.Validator != nil && err != nil {
			goto getAnswer
		}
		// Load the answer
		WriteAnswer(v, q.Name, answer)
	}

	return nil
}
