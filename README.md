# Prompter

This is a simple command line based user interface written in Go.

Heavily insprired by [Survey](https://github.com/AlecAivazis/survey).
# Table of Contents
1. [Quick Start](#quick-start)
1. [Usage](#usage)
    - [Forming Quesitons](#forming-questions)
        -   [Validators](#validators)
        - [Multiselect](#multiselect)
    - [Getting Answers](#getting-answers)
        - [Structs](#ex-1-answer-structs)
        - [Maps](#ex-2-answer-maps)

---

<p align="center">
<img src="https://github.com/AB0529/prompter/raw/main/Showcase-Image.png" data-canonical-src="https://github.com/AB0529/prompter/raw/main/Showcase-Image.png" width="512" />
</p>

# Quick Start
Download the package
```sh
go get github.com/AB0529/prompter
```
```go
package main

import (
	"fmt"

	"github.com/AB0529/prompter"
)

func main() {
	questions := []*prompter.Question{
		{
			Message:   "What is your name?",
			Name:      "name",
			Validator: prompter.Required,
		},
		{
			Message:     "What color?",
			Name:        "color",
			Multiselect: &prompter.Multiselect{"Red", "Green", "Blue", "Purple"},
		},
		{
			Message:   "What is your age?",
			Name:      "age",
			Validator: prompter.IsNumeric,
		},
	}
	answers := struct {
		Name  string
		Color string
		Age   int
	}{}

	err := prompter.Ask(&prompter.Prompt{Questions: questions}, &answers)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Your name is %s, you are %d old, and your favourite color is %s!\n", answers.Name, answers.Age, answers.Color)
}
```

# Usage
## Forming Questions
To ask questions, the prompter will need a slice of pointer questions to ask. Like so:
```go
questions := []*prompter.Question{
    {
        Message:   "What is your name?",
        Name:      "name",
        Validator: prompter.Required,
    },
}
```
A `Question` type has fields you can provide.
- `Message`
    - This is the message that will be prompted to the user.
- `Name`
    - This will be used to identify your question with the answers.
- `Validator`
    - An optional field passed into the question.
    - See [Validators](#Validators).
- `Multiselect`
    - An optional field passed into the question.
    - See [Multiselect](#Multiselect).

## Validators
A validator is a function that will be called and tested on an input. If it fails, the question will simply be asked again until it passes.
### Prebuilt Validators
- `IsNumeric` - makes sure an input is numeric
- `Required` - makes sure an input is not empty

### Constructing a Validator
To construct a validator, simply make a function **exactly** like so:
```go
func MyValidator(input interface{}) error {
    if input != "69" {
        // FAIL
        return errors.New("Value is not a funny number")
    }
    // PASS
    return nil
}
```
In order for a validator to pass, it must return `nil`. For it to fail, it must return an `error` type. 

## Multiselect
Multisect will prompt the user for already existing options, in which they select one by moving up and down using the arrow keys.

The options provided must be a **slice of strings**.

### Example
```go
questions := []*prompter.Question{
    {
        Message: "What color?",
        Name: "color",
        Multiselect: &prompter.Multiselect{"Red", "Green", "Blue", "Purple"},
    },
}
```

---

## Getting Answers
To get usable answers after the questions are asked, you need to call the `prompter.Ask()` function. This function accepts **two parameters**: a **Prompt** struct with the questions in it, and a pointer to the data structure you want the answrs to be formated in.

### Ex. 1) Answer Structs
```go
answers := struct {
    Name  string
    Color string
    Age   int
}{}
err := prompter.Ask(&prompter.Prompt{Questions: questions}, &answers)

fmt.Println(answers.Name)
```
### Ex. 2) Answer Maps
```go
answers := map[string]interface{}{}
err := prompter.Ask(&prompter.Prompt{Questions: questions}, &answers)

fmt.Println(answers["name"])
```