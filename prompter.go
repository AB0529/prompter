package prompter

import (
	"bufio"
	"fmt"
	"os"
)

// Validator runs a function passed into it against the input of a prompt
type Validator func(v interface{}) error

// Input type in which you are asked a prompt, and you provide a response
type Input struct {
	// Name will be later used to  find the answer
	Name string
	// Message the message to print
	Message string
	// Validators slice of validator types, runs response against all of these
	Validators []Validator
}

// Password same as input type, but hides your typing while typing
type Password struct {
	// Name will be later used to  find the answer
	Name string
	// Message the message to print
	Message string
	// Validators slice of validator types, runs response against all of these
	Validators []Validator
}

// Boolean asks a prompt that can only accept a y(es) or n(o) response
type Boolean struct {
	// Name will be later used to  find the answer
	Name string
	// Message the message to print
	Message string
	// Validators slice of validator types, runs response against all of these
	Validators []Validator
}

// Multiselect type in which you select from options provided
type Multiselect struct {
	// Name will be later used to  find the answer
	Name string
	// Message the message to print
	Message string
	// Options the options to select from
	Options []string
	// Validators slice of validator types, runs response against all of these
	Validators []Validator
}

// Prompt asks all the prompts provided and gets the responses
type Prompt struct {
	// Types the prompt types to ask and get response from
	Types []interface{}
}

// Ask will ask all the prompts provided and gather the response
func Ask(p *Prompt, v interface{}) error {
	scanner := bufio.NewScanner(os.Stdin)

	// Loop through each type and handle it respectibly
	for _, t := range p.Types {
		switch t.(type) {
		// Input type
		case *Input:
			resp, err := Inputer(t.(*Input), scanner)
			err = WriteAnswer(v, t.(*Input).Name, resp)
			if err != nil {
				panic(err)
			}
			continue
		// Boolean
		case *Boolean:
			resp, err := Booleaner(t.(*Boolean), scanner)
			err = WriteAnswer(v, t.(*Boolean).Name, resp)
			if err != nil {
				panic(err)
			}
			continue
		// Password
		case *Password:
			resp, err := Passworder(t.(*Password))
			err = WriteAnswer(v, t.(*Password).Name, resp)
			if err != nil {
				panic(err)
			}
			continue
		// Multiselect
		case *Multiselect:
			resp, err := Multiselecter(t.(*Multiselect))
			err = WriteAnswer(v, t.(*Multiselect).Name, resp)
			if err != nil {
				panic(err)
			}
			continue
		default:
			panic(fmt.Sprintf("%T is not a pointer or a valid type", t))
		}
	}

	return nil
}
