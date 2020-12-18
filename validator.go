package prompter

import (
	"errors"
	"strconv"
)

// Required does not allow an empty value
func Required(val interface{}) error {
	if val == "" {
		return errors.New("Value must be required")
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
