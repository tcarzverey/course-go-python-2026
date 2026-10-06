package handler

import "context"

type UsersDB interface {
	UpdateBalance(ctx context.Context, userID, balance int64) error
}
