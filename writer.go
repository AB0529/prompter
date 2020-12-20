package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Settable allow for configuration when assigning answers
type Settable interface {
	WriteAnswer(field string, value interface{}) error
}

// WriteAnswer writes the incomming answers to the value
func WriteAnswer(t interface{}, name string, v interface{}) (err error) {
	// Custom field type
	if s, ok := t.(Settable); ok {
		return s.WriteAnswer(name, v)
	}

	// The target to write to
	target := reflect.ValueOf(t)
	// Value to write from
	value := reflect.ValueOf(v)

	// Make sure target is pointer
	if target.Kind() != reflect.Ptr {
		return errors.New("You must pass a pointer as the target")
	}
	elem := target.Elem()

	switch elem.Kind() {
	case reflect.Struct:
		// Get the name of the field that matches the string we  were given
		fieldIndex, err := findFieldIndex(elem, name)
		if err != nil {
			return err
		}
		field := elem.Field(fieldIndex)
		// Handle references to the Settable interface aswell
		if s, ok := field.Interface().(Settable); ok {
			return s.WriteAnswer(name, v)
		}
		if field.CanAddr() {
			if s, ok := field.Addr().Interface().(Settable); ok {
				// Use the interface method
				return s.WriteAnswer(name, v)
			}
		}

		// Copy the value over to the normal struct
		return copy(field, value)

	case reflect.Map:
		mapType := reflect.TypeOf(t).Elem()
		if mapType.Key().Kind() != reflect.String {
			return errors.New("answer maps key must be of type string")
		}

		if mapType.Elem().Kind() != reflect.Interface {
			return errors.New("answer maps must be of type map[string]interface")
		}
		mt := *t.(*map[string]interface{})
		mt[name] = value.Interface()
		return nil
	}
	// Otherwise just copy the value to the target
	return copy(elem, value)
}

type errFieldNotMatch struct {
	questionName string
}

func (err errFieldNotMatch) Error() string {
	return fmt.Sprintf("could not find field matching %v", err.questionName)
}

func (err errFieldNotMatch) Is(target error) bool {
	if target != nil {
		if name, ok := IsFieldNotMatch(target); ok {
			// If have a filled questionName then perform "deeper" comparison.
			return name == "" || err.questionName == "" || name == err.questionName
		}
	}

	return false
}

// IsFieldNotMatch handle fields not matching
func IsFieldNotMatch(err error) (string, bool) {
	if err != nil {
		if v, ok := err.(errFieldNotMatch); ok {
			return v.questionName, true
		}
	}

	return "", false
}

func findFieldIndex(s reflect.Value, name string) (int, error) {
	// The type of the value
	sType := s.Type()

	// Then look for matching names
	for i := 0; i < sType.NumField(); i++ {
		// The field we are current scanning
		field := sType.Field(i)

		// If the name of the field matches what we're looking for
		if strings.ToLower(field.Name) == strings.ToLower(name) {
			return i, nil
		}
	}

	return -1, errFieldNotMatch{name}
}

// Write takes a value and copies it to the target
func copy(t reflect.Value, v reflect.Value) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// If we paniced with an error
			if _, ok := r.(error); ok {
				// Cast the result to an error object
				err = r.(error)
			} else if _, ok := r.(string); ok {
				// Otherwise we could have paniced with a string so wrap it in an error
				err = errors.New(r.(string))
			}
		}
	}()

	// If we are copying from a string result to something else
	if v.Kind() == reflect.String && v.Type() != t.Type() {
		var castVal interface{}
		var casterr error
		vString := v.Interface().(string)

		switch t.Kind() {
		case reflect.Bool:
			castVal, casterr = strconv.ParseBool(vString)
		case reflect.Int:
			castVal, casterr = strconv.Atoi(vString)
		case reflect.Int8:
			var val64 int64
			val64, casterr = strconv.ParseInt(vString, 10, 8)
			if casterr == nil {
				castVal = int8(val64)
			}
		case reflect.Int16:
			var val64 int64
			val64, casterr = strconv.ParseInt(vString, 10, 16)
			if casterr == nil {
				castVal = int16(val64)
			}
		case reflect.Int32:
			var val64 int64
			val64, casterr = strconv.ParseInt(vString, 10, 32)
			if casterr == nil {
				castVal = int32(val64)
			}
		case reflect.Int64:
			if t.Type() == reflect.TypeOf(time.Duration(0)) {
				castVal, casterr = time.ParseDuration(vString)
			} else {
				castVal, casterr = strconv.ParseInt(vString, 10, 64)
			}
		case reflect.Uint:
			var val64 uint64
			val64, casterr = strconv.ParseUint(vString, 10, 8)
			if casterr == nil {
				castVal = uint(val64)
			}
		case reflect.Uint8:
			var val64 uint64
			val64, casterr = strconv.ParseUint(vString, 10, 8)
			if casterr == nil {
				castVal = uint8(val64)
			}
		case reflect.Uint16:
			var val64 uint64
			val64, casterr = strconv.ParseUint(vString, 10, 16)
			if casterr == nil {
				castVal = uint16(val64)
			}
		case reflect.Uint32:
			var val64 uint64
			val64, casterr = strconv.ParseUint(vString, 10, 32)
			if casterr == nil {
				castVal = uint32(val64)
			}
		case reflect.Uint64:
			castVal, casterr = strconv.ParseUint(vString, 10, 64)
		case reflect.Float32:
			var val64 float64
			val64, casterr = strconv.ParseFloat(vString, 32)
			if casterr == nil {
				castVal = float32(val64)
			}
		case reflect.Float64:
			castVal, casterr = strconv.ParseFloat(vString, 64)
		default:
			return fmt.Errorf("Unable to convert from string to type %s", t.Kind())
		}

		if casterr != nil {
			return casterr
		}

		t.Set(reflect.ValueOf(castVal))
		return
	}

	t.Set(v)

	return
}
