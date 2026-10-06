// В этом файле описаны кастомные ошибки, которые вам надо будет использовать, в этом файле ничего не надо менять

package errors

import "fmt"

// RetryableError ошибка, при возникновении которой мы должны n раз перезапуститься
type RetryableError struct {
	retriesCount int
}

func NewRetryableError(retriesCount int) *RetryableError {
	return &RetryableError{retriesCount: retriesCount}
}

func (r *RetryableError) Error() string {
	return fmt.Sprintf("retriable error")
}

func (r *RetryableError) RetryCount() int {
	return r.retriesCount
}

// NotFoundError ошибка, в случае если не нашли объект
type NotFoundError struct {
	entityID int
}

func NewNotFoundError(entityID int) *NotFoundError {
	return &NotFoundError{entityID: entityID}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found, id=%v", e.entityID)
}

// проверяем что правильно имплементировали интерфейсы
var _ error = (*RetryableError)(nil)
var _ error = (*NotFoundError)(nil)
