package handler

import (
	"context"
	"errors"

	handler_errors "github.com/tcarzverey/course-go-python/homeworks/hw1/handler/errors"
)

type Handler struct {
	db UsersDB
}

func NewHandler(db UsersDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) UpdateUserBalance(ctx context.Context, userID, balance int64) error {
	err := h.db.UpdateBalance(ctx, userID, balance)
	if err == nil {
		return nil
	}

	if retryable, ok := errors.AsType[*handler_errors.RetryableError](err); ok {
		for range retryable.RetryCount() {
			err = h.db.UpdateBalance(ctx, userID, balance)
			if err == nil {
				return nil
			}
		}

		return err
	}

	if _, ok := errors.AsType[*handler_errors.NotFoundError](err); ok {
		return handler_errors.NewAdditionalMessageError(err, "I can write here everything I want")
	}

	if _, ok := errors.AsType[*handler_errors.AdditionalMessageError](err); ok {
		return err
	}

	panic(err)
}
