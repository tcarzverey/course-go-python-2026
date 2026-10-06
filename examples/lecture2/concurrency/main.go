package concurrency

import (
	"sync"
)

func move(wg *sync.WaitGroup, from <-chan int, to chan int) {
	defer wg.Done()

	for el := range from {
		to <- el
	}
}

func merge(chns ...<-chan int) <-chan int {
	res := make(chan int)

	wg := sync.WaitGroup{}

	for _, ch := range chns {
		wg.Add(1)
		go move(&wg, ch, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
