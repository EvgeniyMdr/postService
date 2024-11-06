package errors

import "fmt"

type CustomError struct {
	Err    error
	Code   int
	Source string
}

func New(message string, code int, source string) *CustomError {
	return &CustomError{
		Code:   code,
		Source: source,
	}
}

func Wrap(err error, code int, source string) *CustomError {
	return &CustomError{
		Err:    err,
		Code:   code,
		Source: source,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s]: %v", e.Source, e.Err)
}
