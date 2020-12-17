package prompter

import (
	"errors"
	"reflect"
	"strconv"
)

// Required does not allow an empty value
func Required(val interface{}) error {
	// The reflect value of the result
	value := reflect.ValueOf(val)

	// If the value passed in is the zero value of the appropriate type
	if value.IsZero() && value.Kind() != reflect.Bool {
		return errors.New("Value is required")
	}
	return nil
}

// IsNumeric makes sure the value is numeric
func IsNumeric(val interface{}) error {
	_, err := strconv.Atoi(val.(string))

	// Handle non numeric
	if err != nil {
		return errors.New("Value is not numeric")
	}

	return nil
}
