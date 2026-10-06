package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tcarzverey/course-go-python/homeworks/hw1/handler/errors"
)

type dbStub struct {
	res error
}

func (d *dbStub) UpdateBalance(_ context.Context, _, _ int64) error {
	return d.res
}

func TestHandler_UpdateUserBalanceNoError(t *testing.T) {
	h := NewHandler(&dbStub{})
	userID, balance := int64(123), int64(999)

	err := h.UpdateUserBalance(context.Background(), userID, balance)
	require.NoError(t, err)
}

func TestHandler_UpdateUserBalance_RetryableError(t *testing.T) {
	retryErr := errors.NewRetryableError(2)
	wrappedErr := fmt.Errorf("retryErr: %w", retryErr)
	h := NewHandler(&dbStub{res: wrappedErr})
	userID, balance := int64(1), int64(100)

	err := h.UpdateUserBalance(context.Background(), userID, balance)
	require.Error(t, err)
	require.ErrorIs(t, err, retryErr)
}

func TestHandler_UpdateUserBalance_NotFoundError(t *testing.T) {
	notFoundErr := errors.NewNotFoundError(123)
	wrappedErr := fmt.Errorf("notFoundErr: %w", notFoundErr)
	h := NewHandler(&dbStub{res: wrappedErr})
	userID, balance := int64(2), int64(200)

	err := h.UpdateUserBalance(context.Background(), userID, balance)
	require.Error(t, err)
	require.ErrorIs(t, err, notFoundErr)

	errI := any(err)
	_, ok := errI.(*errors.AdditionalMessageError)
	require.True(t, ok, "error should be wrapped in AdditionalMessageError")
	// Проверяем, что сообщение содержит информацию о том, что пользователь не найден
	require.Contains(t, err.Error(), "not found")
}

func TestHandler_UpdateUserBalance_AdditionalMessageError(t *testing.T) {
	additionalErr := errors.NewAdditionalMessageError(fmt.Errorf("not found"), "not found")
	wrappedErr := fmt.Errorf("additionalErr: %w", additionalErr)
	h := NewHandler(&dbStub{res: wrappedErr})
	userID, balance := int64(2), int64(200)

	err := h.UpdateUserBalance(context.Background(), userID, balance)
	require.Error(t, err)
	require.ErrorIs(t, err, additionalErr)
	require.Equal(t, wrappedErr, err)
}

func TestHandler_UpdateUserBalance_UnknownError_Panics(t *testing.T) {
	unknownErr := fmt.Errorf("some unknown error")
	h := NewHandler(&dbStub{res: unknownErr})
	userID, balance := int64(3), int64(300)

	require.Panics(t, func() {
		_ = h.UpdateUserBalance(context.Background(), userID, balance)
	})
}
