package hw09structvalidator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type intValidator struct {
	name      string
	condition string
	field     string
	value     int
}

func intHandlers(tag, field string, value int) []intValidator {
	validatorsRaw := strings.Split(tag, delimiterMulti)
	validators := make([]intValidator, 0)

	for _, valRaw := range validatorsRaw {
		val := strings.Split(valRaw, delimiter)
		if len(val) != 2 {
			log.Printf("invalid value for tag %s\n", tag)
			continue
		}
		sVal := intValidator{
			name:      val[0],
			condition: val[1],
			field:     field,
			value:     value,
		}
		validators = append(validators, sVal)
	}
	return validators
}

func (iv intValidator) Validate() *ValidationError {
	switch iv.name {
	case "min":
		cond, err := strconv.Atoi(iv.condition)
		if err != nil {
			log.Println("min value is not int")
			return nil
		}
		if iv.value < cond {
			return &ValidationError{
				Field: iv.field,
				Err:   fmt.Errorf("%w: value %d, condition %d", ErrMin, iv.value, cond),
			}
		}

	case "max":
		cond, err := strconv.Atoi(iv.condition)
		if err != nil {
			log.Println("max value is not int")
			return nil
		}
		if iv.value > cond {
			return &ValidationError{
				Field: iv.field,
				Err:   fmt.Errorf("%w: value %d, condition %d", ErrMax, iv.value, cond),
			}
		}

	case "in":
		err := iv.in()
		if err != nil {
			return err
		}
	default:
		log.Printf("unknown validator's name %s", iv.name)
	}
	return nil
}

func (iv intValidator) in() *ValidationError {
	set := strings.Split(iv.condition, ",")
	for _, val := range set {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Println("set's value is not int")
			return nil
		}
		if iv.value == intVal {
			return nil
		}
	}
	return &ValidationError{
		Field: iv.field,
		Err:   fmt.Errorf("%w: value %d, set %v", ErrNotIn, iv.value, set),
	}
}
