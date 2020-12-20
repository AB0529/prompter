package prompter

import "github.com/gookit/color"

var (
	// Red color red
	Red = color.Red
	// Blue color blue
	Blue = color.Blue
	// Cyan color cyan
	Cyan = color.Cyan
	// Purple color purple
	Purple = color.Magenta
	// Green color green
	Green = color.LightGreen
	// Yellow color yellow
	Yellow = color.Yellow
	// White color white
	White = color.White
)

var (
	// InputChar color of the "> "
	InputChar = Red
	// Title the title of questions
	Title = Cyan
	// ValidateError value error from validator
	ValidateError = Red
	// MultiselectOptions options for multiselect
	MultiselectOptions = Yellow
	// MultiselectAnswer answer selected
	MultiselectAnswer = Green
	// BooleanPrompt color for boolean [Y/N]
	BooleanPrompt = Yellow
)
