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

// Validator validates function passed into it, MUST RETURN ERROR
type Validator func(ans interface{}) error

// Question structure for a question that'll be asked
type Question struct {
	Message   string      `json:"message" binding:"required"`
	Name      string      `json:"name" binding:"required"`
	Validator Validator   `json:"validator,omitempty"`
	Type      interface{} `json:"type,omitempty"`
}

// Prompt the promt which will ask the questions provided and get the answers
type Prompt struct {
	Questions []*Question
}

func answerFunc(q *Question, scanner *bufio.Scanner, err error) string {
	// Print the question
	if err != nil {
		fmt.Print(Purple(q.Message) + "\n" + fmt.Sprintf("[%s]", White(err.Error())) + Red(" > "))
	} else {
		fmt.Print(Purple(q.Message) + "\n" + Red("> "))
	}
	scanner.Scan()

	return scanner.Text()
}

// Ask will actually ask the questions and get the answers
func Ask(p *Prompt, v interface{}) error {
	var err error
	var answer string
	scanner := bufio.NewScanner(os.Stdin)

	for _, q := range p.Questions {
		// Handle Multiselect
		switch q.Type.(type) {
		case Multiselect:
			// Print question
			fmt.Println(Purple(q.Message))
			// Print options
			for _, ms := range q.Type.(Multiselect) {
				fmt.Println(Red("> ") + Yellow(ms))
			}
			// Get the answer
			WriteAnswer(v, q.Name, Multiselector(q))
			continue
		case Password:
			err = q.Validator(answer)

			// Handle validator
			if q.Validator != nil && err != nil {
				goto getAnswer
			}
			// Load the answer
			answer = PasswordSelector(q)
			WriteAnswer(v, q.Name, answer)
			fmt.Println()
			continue
		}

		// Handle validator
		if q.Validator != nil && err != nil {
			goto getAnswer
		}

	getAnswer:
		answer = answerFunc(q, scanner, err)
		// Load the answer
		WriteAnswer(v, q.Name, answer)
	}

	return nil
}
