package hw09structvalidator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	tagName        = "validate"
	delimiter      = ":"
	delimiterMulti = "|"
)

var (
	ErrLength    = errors.New("length is invalid")
	ErrNotStruct = errors.New("struct is invalid")
	ErrRegexp    = errors.New("regexp is invalid")
	ErrNotIn     = errors.New("in is element")
	ErrMin       = errors.New("min is invalid")
	ErrMax       = errors.New("max is invalid")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	s := strings.Builder{}

	for _, e := range v {
		s.WriteString(fmt.Sprintf("field: %q - %s\n", e.Field, e.Err.Error()))
	}

	return s.String()
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors
	value := reflect.ValueOf(v)

	if value.Kind() != reflect.Struct {
		validationErrors = append(validationErrors, ValidationError{
			Field: "",
			Err:   ErrNotStruct,
		})
		return validationErrors
	}

	fieldType := value.Type()

	for i := 0; i < fieldType.NumField(); i++ {
		fieldValue := fieldType.Field(i)
		tag := fieldValue.Tag.Get(tagName)

		if len(tag) == 0 {
			continue
		}

		validateValue := value.Field(i)

		if !validateValue.CanInterface() {
			continue
		}
		validationErrors = validateSwitcher(tag, fieldValue.Name, validateValue, validationErrors)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateSwitcher(tag, field string, value reflect.Value, errors ValidationErrors) ValidationErrors {
	switch ref := value.Kind(); ref {
	case reflect.String:
		validators := stringHandlers(tag, field, value.String())
		for _, validator := range validators {
			err := validator.validate()
			if err != nil {
				errors = append(errors, *err)
			}
		}
		return errors

	case reflect.Int:
		validators := intHandlers(tag, field, int(value.Int()))
		for _, validator := range validators {
			err := validator.Validate()
			if err != nil {
				errors = append(errors, *err)
			}
		}
		return errors

	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			elem := value.Index(i)
			errors = append(errors, validateSwitcher(tag, field, elem, ValidationErrors{})...)
		}
		return errors
	default:
		log.Println("unsupported type")
		return nil
	}
}
