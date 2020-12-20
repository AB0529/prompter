package main

import (
	"bufio"
	"errors"
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

// Booleaner will handle the boolean type and get the response
func Booleaner(t *Boolean, scanner *bufio.Scanner) (bool, error) {
	var err error
PrintQuestion:
	// Print question
	if err != nil {
		fmt.Printf(fmt.Sprintf("[%s] ", BooleanPrompt.Sprint("Y/N")))
		printQuestionWithError(t.Message, err)
	} else {
		fmt.Printf(fmt.Sprintf("[%s] ", BooleanPrompt.Sprint("Y/N")))
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

	if resp == "" {
		return true, nil
	}

	switch string(resp[0]) {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		err = errors.New("Value unknown")
		goto PrintQuestion
	}

}
