package main

import (
	"fmt"

	"github.com/go-playground/validator/v10/non-standard/validators"
)

type Error interface {
	Error() string
}

type Errorer interface {
	Error() string
}

type CustomError struct {
}

func (e *CustomError) Error() string {
	return ""
}

func (e *CustomError) String() string {
	return "This is a string"
}

func main() {
	var a any = 1
	var _ error = &CustomError{}
	var _ fmt.Stringer = &CustomError{}

	var customError Errorer
	if converterdError, ok := (customError).(error); ok {

	}

	validators.NotBlank()
}
