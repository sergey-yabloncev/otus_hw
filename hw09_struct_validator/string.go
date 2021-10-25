package hw09structvalidator

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const delimerIn = ","

type stringValidator struct {
	name      string
	condition string
	field     string
	value     string
}

func stringHandlers(tag, field, value string) []stringValidator {
	validatorsRaw := strings.Split(tag, delimerMulti)
	validators := make([]stringValidator, 0)

	for _, valRaw := range validatorsRaw {
		val := strings.Split(valRaw, delimer)
		if len(val) != 2 {
			log.Printf("invalid value for tag %s\n", tag)
			continue
		}
		sVal := stringValidator{
			name:      val[0],
			condition: val[1],
			field:     field,
			value:     value,
		}
		validators = append(validators, sVal)
	}
	return validators
}

func (sv stringValidator) validate() *ValidationError {
	switch sv.name {
	case "len":
		cond, err := strconv.Atoi(sv.condition)
		if err != nil {
			log.Println("len value is not int")
			return nil
		}
		l := len(sv.value)
		if l != cond {
			return &ValidationError{
				Field: sv.field,
				Err:   fmt.Errorf("%w: need %d, now %d", ErrLength, cond, l),
			}
		}

	case "regexp":
		err := sv.regex()
		if err != nil {
			return err
		}
	case "in":
		set := strings.Split(sv.condition, delimerIn)
		for _, e := range set {
			if sv.value == e {
				return nil
			}
		}
		return &ValidationError{
			Field: sv.field,
			Err:   fmt.Errorf("%w: value %s, set %v", ErrNotIn, sv.value, set),
		}

	default:
		log.Printf("unknown validator's name %s", sv.name)
	}
	return nil
}

func (sv stringValidator) regex() *ValidationError {
	matched, err := regexp.MatchString(sv.condition, sv.value)
	if err != nil {
		return &ValidationError{Field: sv.field, Err: err}
	}
	if !matched {
		return &ValidationError{
			Field: sv.field,
			Err:   fmt.Errorf("%w: %s", ErrRegexp, sv.condition),
		}
	}
	return nil
}
