package urls

import (
	"context"
	"log"
	"net/http"
	"sync"
)

const workersCount = 8

type ResponseCodeAggregator struct {
	client HttpClient
}

func NewURLAggregator(client HttpClient) *ResponseCodeAggregator {
	return &ResponseCodeAggregator{client: client}
}

// Aggregate запускает агрегацию асинхронно и сразу возвращает результат.
func (u *ResponseCodeAggregator) Aggregate(ctx context.Context, urls <-chan string) (AggregationResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	res := newAggregationResult()

	var wg sync.WaitGroup
	for range workersCount {
		wg.Go(func() {
			u.worker(ctx, urls, res)
		})
	}

	go func() {
		wg.Wait()
		res.finish()
	}()

	return res, nil
}

// worker обрабатывает url из канала до его закрытия или отмены контекста.
func (u *ResponseCodeAggregator) worker(ctx context.Context, urls <-chan string, res *aggregationResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-urls:
			if !ok {
				return
			}
			u.process(ctx, url, res)
		}
	}
}

type responseReply struct {
	resp *http.Response
	err  error
}

// process выполняет GET-запрос, не дожидаясь ответа при отмене контекста.
func (u *ResponseCodeAggregator) process(ctx context.Context, url string, res *aggregationResult) {
	replies := make(chan responseReply)
	go func() {
		resp, err := u.client.Get(url)
		reply := responseReply{resp: resp, err: err}
		select {
		case replies <- reply:
		case <-ctx.Done():
			if reply.resp != nil {
				_ = reply.resp.Body.Close()
			}
		}
	}()

	select {
	case <-ctx.Done():
		return
	case reply := <-replies:
		if reply.err != nil {
			log.Printf("get %q: %v", url, reply.err)

			return
		}
		defer reply.resp.Body.Close()
		res.add(reply.resp.StatusCode)
	}
}
