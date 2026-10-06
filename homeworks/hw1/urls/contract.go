package urls

import "net/http"

// AggregationResult интерфейс, под который нужно написать имплементацию в result.go
type AggregationResult interface {
	// GetResponsesCount возвращает количество запросов с заданным кодом ответа
	GetResponsesCount(code int) int
	// GetResult возвращает мапу вида код ответа => количество запросов с этим кодом
	GetResult() map[int]int
	// Done возвращает true, если агрегация завершена
	Done() bool
}

type HttpClient interface {
	// Get выполняет GET запрос по указанному url
	Get(url string) (resp *http.Response, err error)
}
