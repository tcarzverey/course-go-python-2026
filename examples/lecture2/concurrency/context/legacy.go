package context

import (
	"context"
	"time"
)

func goodCall(ctx context.Context, userID int64) (int64, error) {
	var httpCaller interface {
		Get(context.Context, int64) (int64, error)
	}

	return httpCaller.Get(ctx, userID)

}

func legacyCall(userID int64) (int64, error) {
	time.Sleep(time.Second * 5)
	return userID + 100, nil
}
