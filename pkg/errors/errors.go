package errors

import "fmt"

type AppError struct {
	Message string
	Err     error
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func Wrap(err error, message string) AppError {
	return AppError{
		Message: message,
		Err:     err,
	}
}
