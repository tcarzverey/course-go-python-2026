package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdditionalMessageErrorErrorMessage(t *testing.T) {
	// если передали формат и ошибку, выводим сначала результат форматирования, потом двоеточие, пробел, и после текст вложенной ошибки
	err := NewAdditionalMessageError(errors.New("wrapped error text"), "got err, id=%v, operation=%s", 123, "create")
	require.Error(t, err)
	assert.Equal(t, "got err, id=123, operation=create: wrapped error text", err.Error())
	err = NewAdditionalMessageError(errors.New("err"), "error")
	require.Error(t, err)
	assert.Equal(t, "error: err", err.Error())

	// если строка формата пустая - возвращаем только текст ошибки
	err = NewAdditionalMessageError(errors.New("no additional message"), "")
	require.Error(t, err)
	assert.Equal(t, "no additional message", err.Error())

	// если ошибка nil - возвращаем просто отформатированную строку
	err = NewAdditionalMessageError(nil, "got nil error")
	require.Error(t, err)
	assert.Equal(t, "got nil error", err.Error())

	// если ошибка nil и формат пустой - возвращаем пустую строку
	err = NewAdditionalMessageError(nil, "")
	require.Error(t, err)
	assert.Equal(t, "", err.Error())
}

func TestAdditionalMessageErrorUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := NewAdditionalMessageError(originalErr, "additional context: %s", "test")

	// Проверяем, что unwrap возвращает оригинальную ошибку
	unwrappedErr := errors.Unwrap(wrappedErr)
	require.Equal(t, originalErr, unwrappedErr)

	// Проверяем, что errors.Is работает правильно
	require.True(t, errors.Is(wrappedErr, originalErr))
}

func TestAdditionalMessageErrorNilUnwrap(t *testing.T) {
	wrappedErr := NewAdditionalMessageError(nil, "no original error")

	// Проверяем, что unwrap возвращает nil для nil ошибки
	unwrappedErr := errors.Unwrap(wrappedErr)
	require.Nil(t, unwrappedErr)
}
