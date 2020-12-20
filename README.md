# Prompter

This is a simple command line based user interface written in Go.

Heavily insprired by [Survey](https://github.com/AlecAivazis/survey).
# Table of Contents
1. [Quick Start](#quick-start)
1. [Usage](#usage)
    - [Forming Quesitons](#forming-questions)
        -  [Validators](#validators)
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
	questions := []interface{}{
		&prompter.Input{
			Name:       "name",
			Message:    "What is your name?",
			Validators: []prompter.Validator{prompter.Required},
		},
		&prompter.Boolean{
			Name:    "candy",
			Message: "You like candy?",
		},
		&prompter.Password{
			Name:       "pass",
			Message:    "What's your password?",
			Validators: []prompter.Validator{prompter.Required},
		},
		&prompter.Multiselect{
			Name:    "color",
			Message: "Color",
			Options: []string{"Red", "Blue", "Green"},
		},
	}
	answer := struct {
		Name  string
		Candy bool
		Pass  string
		Color string
	}{}

	prompter.Ask(&prompter.Prompt{Types: questions}, &answer)
	fmt.Println(answer)
```

# Usage
## Forming Questions
To ask questions, you'll need to provide a slice of empty interfaces with the question types like so:
```go
questions := []interface{}{
    &prompter.Input{
        Name:       "name",
        Message:    "What is your name?",
        Validators: []prompter.Validator{prompter.Required},
    },
}
```
There are multiple types, each aimed at a different style of prompts.
- `Types`
    - `Input`
        - Asks for a prompt and gets the response
    - `Password`
        - Hides an input being typed
    - `Boolean`
        -  Prompts for a y(es)/n(o) answer
    - `Multiselect`
        - See [Multiselect](#Multiselect).

Each type has fields you can provide.
- `Message`
    - This is the message that will be prompted to the user.
- `Name`
    - This will be used to identify your question with the answers.
- `Validators`
    - See [Validators](#Validators).

## Validators
A validator is a function that will be called and tested on an input. If it fails, the question will simply be asked again until it passes.
### Prebuilt Validators
- `IsNumeric` - makes sure an input is numeric
- `Required` - makes sure an input is not empty
- `IsURL` - makes sure an input is a valid URL

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
questions := []interface{}{
    &prompter.Multiselect{
        Name:       "color",
        Message:   "What color you like?",
        Options:    []string{"Red", "Green", "Blue"},
        Validators: []prompter.Validator{prompter.Required},
    },
}
```

---

## Getting Answers
To get usable answers after the questions are asked, you need to call the `prompter.Ask()` function. This function accepts **two parameters**: a **Prompt** struct with the questions in it, and a pointer to the data structure you want the answrs to be formated in.

### Ex. 1) Answer Structs
```go
answer := struct {
    Name  string
    Candy bool
    Pass  string
    Color string
}{}
prompter.Ask(&prompter.Prompt{Types: questions}, &answer)

fmt.Println(answers.Name)
```
### Ex. 2) Answer Maps
```go
answers := map[string]interface{}{}
prompter.Ask(&prompter.Prompt{Types: questions}, &answers)

fmt.Println(answers["name"])
```