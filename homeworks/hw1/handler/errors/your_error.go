package errors

import "fmt"

// AdditionalMessageError ошибка, которая добавляет дополнительный контекст к вложенной ошибке.
type AdditionalMessageError struct {
	err error
	msg string
}

// NewAdditionalMessageError создает ошибку
func NewAdditionalMessageError(err error, format string, args ...any) error {
	return &AdditionalMessageError{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
}

// Error формирует итоговое сообщение ошибки.
func (e *AdditionalMessageError) Error() string {
	switch {
	case e.err != nil && e.msg != "":
		return e.msg + ": " + e.err.Error()
	case e.err != nil:
		return e.err.Error()
	default:
		return e.msg
	}
}

// Unwrap отдает вложенную ошибку в error chain.
func (e *AdditionalMessageError) Unwrap() error {
	return e.err
}

var _ error = (*AdditionalMessageError)(nil)
