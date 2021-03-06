package prompter

import (
	"errors"
	"net/url"
	"strconv"
)

// Required does not allow an empty value
func Required(val interface{}) error {
	if val == "" {
		return errors.New("value is required")
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

// IsURL makes sure a value is a valid URL
func IsURL(val interface{}) error {
	_, err := url.ParseRequestURI(val.(string))

	if err != nil {
		return errors.New("Value is invalid URL")
	}

	return nil
}
