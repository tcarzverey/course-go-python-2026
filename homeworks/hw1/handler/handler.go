package handler

import (
	"context"
)

type Handler struct {
	db UsersDB
}

func NewHandler(db UsersDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) UpdateUserBalance(ctx context.Context, userID, balance int64) error {
	err := h.db.UpdateBalance(ctx, userID, balance)
	// TODO: сюда вам нужно добавить логику обработки различных ошибок
	return err
}
