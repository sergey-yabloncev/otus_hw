package hw09structvalidator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in             interface{}
		expectedFields []string
		expectedErrors []error
	}{
		{
			in: User{
				ID:     "12345",
				Age:    29,
				Email:  "sergey@gmail.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
			},
			expectedFields: []string{
				"ID",
			},
			expectedErrors: []error{
				ErrLength,
			},
		},
		{
			in: App{
				Version: "12345",
			},
		},
		{
			in: App{
				Version: "123456",
			},
			expectedFields: []string{
				"Version",
			},
			expectedErrors: []error{
				ErrLength,
			},
		},
		{
			in: Token{},
		},
		{
			in: Response{
				Code: 200,
				Body: "Ok",
			},
		},
		{
			in: Response{
				Code: 310,
				Body: "Ok",
			},
			expectedFields: []string{
				"Code",
			},
			expectedErrors: []error{
				ErrNotIn,
			},
		},
		{
			in: nil,
			expectedFields: []string{
				"",
			},
			expectedErrors: []error{
				ErrNotStruct,
			},
		},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			errs := expandErrors(err)

			require.Len(t, errs, len(tt.expectedErrors))
			if len(errs) > 0 {
				for i, err := range errs {
					require.Equal(t, err.Field, tt.expectedFields[i])
					require.True(t, errors.Is(err.Err, tt.expectedErrors[i]))
				}
			}
		})
	}
}

func expandErrors(err error) ValidationErrors {
	var validationErrors ValidationErrors
	if !errors.As(err, &validationErrors) {
		return nil
	}
	return validationErrors
}
