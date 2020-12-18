package prompter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gookit/color"
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
	Validator []Validator `json:"validator,omitempty"`
	Type      interface{} `json:"type,omitempty"`
}

// Prompt the promt which will ask the questions provided and get the answers
type Prompt struct {
	Questions []*Question
}

func answerFunc(q *Question, scanner *bufio.Scanner, err error) string {
	// Print the question
	if err != nil {
		fmt.Print(color.Cyan.Sprint(q.Message) + "\n" + fmt.Sprintf("[%s]", color.White.Sprint(err.Error())) + color.Red.Sprint(" > "))
	} else {
		fmt.Print(color.Cyan.Sprint(q.Message) + "\n" + color.Red.Sprint("> "))
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
			fmt.Println(color.Cyan.Sprint(q.Message))
			// Print options
			for _, ms := range q.Type.(Multiselect) {
				fmt.Println(color.Red.Sprint("> ") + color.Yellow.Sprint(ms))
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
		}

	getAnswer:
		answer = answerFunc(q, scanner, err)
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
