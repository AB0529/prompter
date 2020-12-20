package main

import (
	"bufio"
	"fmt"
)

func printQuestion(q string) {
	fmt.Printf("%s\n%s ", Title.Sprint(q), InputChar.Sprint(">"))
}

func printQuestionWithError(q string, err error) {
	fmt.Printf("%s (%s)\n%s ", Title.Sprint(q), ValidateError.Sprint(err.Error()), InputChar.Sprint(">"))
}

// Inputer will handle the input type and get the response
func Inputer(t *Input, scanner *bufio.Scanner) (string, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		printQuestionWithError(t.Message, err)
	} else {
		printQuestion(t.Message)
	}
	scanner.Scan()

	// Get response
	resp := scanner.Text()

	// Handle validators with response
	for _, val := range t.Validators {
		// Run the validator function
		err = val(resp)
		if err != nil {
			goto PrintQuestion
		}
	}

	return resp, nil
}
