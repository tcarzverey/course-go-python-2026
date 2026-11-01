package urls

import (
	"maps"
	"sync"
)

// aggregationResult потокобезопасная реализация AggregationResult.
type aggregationResult struct {
	mu     sync.RWMutex
	counts map[int]int
	doneCh chan struct{}
}

func newAggregationResult() *aggregationResult {
	return &aggregationResult{
		counts: make(map[int]int),
		doneCh: make(chan struct{}),
	}
}

// GetResponsesCount возвращает количество ответов с кодом code.
func (r *aggregationResult) GetResponsesCount(code int) int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.counts[code]
}

// GetResult возвращает копию мапы код ответа => количество запросов.
func (r *aggregationResult) GetResult() map[int]int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return maps.Clone(r.counts)
}

// Done возвращает true, если агрегация завершена.
func (r *aggregationResult) Done() bool {
	select {
	case <-r.doneCh:
		return true
	default:
		return false
	}
}

// add увеличивает счетчик для кода ответа.
func (r *aggregationResult) add(code int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counts[code]++
}

// finish помечает агрегацию завершенной.
func (r *aggregationResult) finish() {
	close(r.doneCh)
}

var _ AggregationResult = (*aggregationResult)(nil)
