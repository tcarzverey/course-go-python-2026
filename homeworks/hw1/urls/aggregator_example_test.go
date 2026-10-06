package urls

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAggregatorSimple(t *testing.T) {
	client := &http.Client{}
	aggregator := NewURLAggregator(client)

	urls := make(chan string)
	go func() {
		urls <- "http://www.google.com"
		urls <- "http://ya.ru"
		close(urls)
	}()

	_, err := aggregator.Aggregate(context.Background(), urls)
	require.NoError(t, err)
}

func TestAggregatorFinalResult(t *testing.T) {
	client := &http.Client{}
	aggregator := NewURLAggregator(client)

	urls := make(chan string)
	go func() {
		urls <- "http://www.google.com"
		urls <- "http://ya.ru"
		close(urls)
	}()

	res, err := aggregator.Aggregate(context.Background(), urls)
	require.NoError(t, err)
	require.NotNil(t, res)
	time.Sleep(time.Second)
	assert.True(t, res.Done())
	assert.Equal(t, 2, res.GetResponsesCount(http.StatusOK))
	assert.Equal(t, 0, res.GetResponsesCount(http.StatusNotFound))
	assert.Equal(t, map[int]int{http.StatusOK: 2}, res.GetResult())
}
